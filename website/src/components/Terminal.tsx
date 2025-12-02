import { useState } from 'react';
import { Copy, Check } from 'lucide-react';
import { Button } from '@/components/ui/button';

export default function Terminal() {
  const [copied, setCopied] = useState(false);
  const command = 'npx doplan@latest init';

  const handleCopy = async () => {
    await navigator.clipboard.writeText(command);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="relative w-full max-w-2xl mx-auto">
      <div className="rounded-lg bg-black shadow-2xl overflow-hidden border border-white/10">
        {/* Terminal Header */}
        <div className="flex items-center gap-2 px-4 py-3 bg-gradient-to-r from-gray-800 to-gray-900 border-b border-white/10">
          <div className="flex gap-2">
            <div className="w-3 h-3 rounded-full bg-red-500" />
            <div className="w-3 h-3 rounded-full bg-yellow-500" />
            <div className="w-3 h-3 rounded-full bg-green-500" />
          </div>
          <div className="flex-1 text-center text-xs text-gray-400 font-mono">
            terminal
          </div>
        </div>

        {/* Terminal Body */}
        <div className="relative p-6 font-mono text-sm">
          <div className="flex items-center gap-3">
            <span className="text-green-400">$</span>
            <span className="text-gray-100 flex-1">{command}</span>
            <Button
              variant="ghost"
              size="icon"
              onClick={handleCopy}
              className="h-8 w-8 text-gray-400 hover:text-white hover:bg-white/10"
            >
              {copied ? (
                <Check className="h-4 w-4 text-green-400" />
              ) : (
                <Copy className="h-4 w-4" />
              )}
            </Button>
          </div>
          
          {copied && (
            <div className="absolute top-2 right-2 text-xs text-green-400 animate-in fade-in slide-in-from-top-1">
              Copied!
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
