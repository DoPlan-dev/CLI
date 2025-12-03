# Vercel Build Diagnostic Guide

## 🔍 What to Look For in Build Logs

### Common Failure Points (in order of occurrence):

1. **Installation Phase** (0-2 minutes)
   - Look for: `pnpm install` errors, lockfile conflicts, dependency resolution failures
   - Error patterns: `ERR_PNPM_*`, `ENOENT`, `EACCES`, version conflicts

2. **TypeScript Check** (2-3 minutes)
   - Look for: `tsc --noEmit` errors, type errors, missing type definitions
   - Error patterns: `TS2307`, `TS2322`, `TS2345`, compilation errors

3. **Vite Build** (3-5 minutes)
   - Look for: Module resolution errors, import errors, build timeouts
   - Error patterns: `Cannot find module`, `Failed to resolve`, build hangs

4. **Post-Build** (5+ minutes)
   - Look for: Deployment errors, file not found, routing issues
   - Error patterns: `ENOENT`, `404`, deployment failures

## 🚨 Most Common Issues & Fixes

### Issue 1: Build Hangs After "done" (Installation Complete)

**Symptoms:**
```
Lockfile is up to date, resolution step is skipped
[... packages installed ...]
done
[Build hangs here - no output for 5+ minutes]
```

**Root Causes:**
1. **TypeScript compilation hanging** - `tsc --noEmit` might be stuck
2. **Vite build not starting** - Missing or broken vite.config.ts
3. **Memory exhaustion** - Build process running out of memory
4. **Infinite loop** - Circular dependencies or recursive imports

**Fixes:**

#### Fix A: Skip TypeScript Check Temporarily
```json
// vercel.json
{
  "buildCommand": "pnpm run build:skip-check"
}
```

#### Fix B: Increase Memory/Timeout
```json
// vercel.json
{
  "buildCommand": "pnpm run build",
  "installCommand": "pnpm install --frozen-lockfile",
  "framework": "vite",
  "outputDirectory": "dist",
  "functions": {
    "website/src/**/*.ts": {
      "maxDuration": 60
    }
  }
}
```

#### Fix C: Add Build Timeout Protection
```json
// package.json - add this script
{
  "scripts": {
    "build:vercel": "timeout 300 pnpm run build || pnpm run build:skip-check"
  }
}
```

### Issue 2: "Cannot find module" Errors

**Symptoms:**
```
Error: Cannot find module '@/components/...'
Error: Failed to resolve import
```

**Fix:** Ensure path aliases are configured in both `tsconfig.json` and `vite.config.ts` ✅ (Already fixed)

### Issue 3: Memory/Timeout Errors

**Symptoms:**
```
Error: JavaScript heap out of memory
Build exceeded maximum duration
```

**Fixes:**

1. **Optimize Vite Config:**
```typescript
// vite.config.ts - add these optimizations
export default defineConfig({
  build: {
    chunkSizeWarningLimit: 1000,
    rollupOptions: {
      output: {
        manualChunks(id) {
          if (id.includes('node_modules')) {
            if (id.includes('react') || id.includes('react-dom')) {
              return 'react-vendor';
            }
            if (id.includes('@radix-ui')) {
              return 'radix-vendor';
            }
            return 'vendor';
          }
        },
      },
    },
  },
});
```

2. **Add .vercelignore:**
```
node_modules
.git
.cache
dist
*.log
coverage
.vscode
.idea
```

### Issue 4: pnpm Version Mismatch

**Symptoms:**
```
Lockfile is out of sync
pnpm-lock.yaml version mismatch
```

**Fix:** Already using `--frozen-lockfile` ✅

### Issue 5: Monorepo Root Directory

**Symptoms:**
```
Cannot find package.json
Build command failed
```

**Fix:** In Vercel Dashboard → Settings → General:
- Set **Root Directory** to: `website`
- Or add to `vercel.json`:
```json
{
  "rootDirectory": "website"
}
```

## 🎯 Optimized Configuration

### Enhanced vercel.json
```json
{
  "buildCommand": "pnpm run build",
  "installCommand": "pnpm install --frozen-lockfile --prefer-offline",
  "framework": "vite",
  "outputDirectory": "dist",
  "devCommand": "pnpm run dev",
  "regions": ["iad1"],
  "env": {
    "NODE_ENV": "production",
    "NODE_OPTIONS": "--max-old-space-size=4096"
  }
}
```

### Enhanced package.json Scripts
```json
{
  "scripts": {
    "dev": "vite",
    "build": "tsc --noEmit && vite build",
    "build:skip-check": "vite build",
    "build:vercel": "NODE_OPTIONS='--max-old-space-size=4096' pnpm run build",
    "lint": "eslint . --ext ts,tsx --report-unused-disable-directives --max-warnings 0",
    "preview": "vite preview"
  }
}
```

### Enhanced vite.config.ts (Performance)
```typescript
import path from "path";
import { fileURLToPath } from "url";
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  build: {
    outDir: "dist",
    sourcemap: false,
    minify: "esbuild",
    target: "es2020",
    chunkSizeWarningLimit: 1000,
    rollupOptions: {
      output: {
        manualChunks(id) {
          // Split vendor chunks for better caching
          if (id.includes('node_modules')) {
            if (id.includes('react') || id.includes('react-dom')) {
              return 'react-vendor';
            }
            if (id.includes('@radix-ui')) {
              return 'radix-vendor';
            }
            if (id.includes('react-router')) {
              return 'router-vendor';
            }
            return 'vendor';
          }
        },
      },
    },
  },
  // Optimize dependencies
  optimizeDeps: {
    include: ['react', 'react-dom', 'react-router-dom'],
  },
});
```

## 📊 Build Performance Optimization

### 1. Enable Build Caching
Vercel automatically caches:
- `node_modules` (with `--frozen-lockfile`)
- Build outputs (if configured)

### 2. Reduce Bundle Size
- Already using manual chunks ✅
- Consider lazy loading routes
- Remove unused dependencies

### 3. Parallel Build Steps
Vercel runs install and build in sequence, but you can optimize:
```json
// Not applicable for Vite, but good to know
```

## 🔧 Environment Variables Checklist

Check if these are set in Vercel Dashboard → Settings → Environment Variables:

- `NODE_ENV=production` (auto-set by Vercel)
- `SUPABASE_PROJECT_ID` (if using Supabase)
- `SUPABASE_URL` (if using Supabase)
- `SUPABASE_ANON_KEY` (if using Supabase)
- Any other API keys your app needs

## 🐛 Debugging Steps

1. **Get Full Log:**
   - Vercel Dashboard → Deployments → Click failed deployment → View Build Logs
   - Copy the ENTIRE log, especially the error message

2. **Test Locally:**
   ```bash
   cd website
   pnpm install --frozen-lockfile
   pnpm run build
   ```

3. **Check for TypeScript Errors:**
   ```bash
   pnpm exec tsc --noEmit
   ```

4. **Test Vite Build Only:**
   ```bash
   pnpm run build:skip-check
   ```

5. **Check Bundle Size:**
   ```bash
   pnpm run build
   du -sh dist/
   ```

## 📝 What to Share for Full Diagnosis

When sharing the build log, include:

1. **Full log from start to error** (or timeout)
2. **The exact error message** (if any)
3. **Where it fails** (install/build/deploy phase)
4. **Build duration** before failure
5. **Any warnings** before the error

## 🚀 Quick Wins for Faster Builds

1. ✅ Use `--frozen-lockfile` (already done)
2. ✅ Enable framework detection (already done)
3. ✅ Optimize chunk splitting (already done)
4. ⚠️ Consider skipping TypeScript check in CI (if types are checked in PRs)
5. ⚠️ Use `--prefer-offline` for faster installs
6. ⚠️ Set `NODE_OPTIONS` for memory management

## ⚡ Expected Build Times

- **Cold Build** (no cache): 4-6 minutes
- **Warm Build** (with cache): 2-4 minutes
- **Install**: 30-60 seconds
- **Type Check**: 10-30 seconds
- **Vite Build**: 1-3 minutes
- **Deploy**: 30-60 seconds

If builds exceed 10 minutes, there's likely an issue.

