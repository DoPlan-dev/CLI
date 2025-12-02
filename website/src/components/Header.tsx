import { Moon, Sun, Github, BookOpen, Terminal } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { useTheme } from '@/hooks/useTheme';
import { Link } from 'react-router-dom';

export default function Header() {
  const { actualTheme, setTheme } = useTheme();

  const toggleTheme = () => {
    setTheme(actualTheme === 'dark' ? 'light' : 'dark');
  };

  return (
    <header className="fixed top-0 left-0 right-0 z-50 border-b border-border/40 bg-background/80 backdrop-blur-lg">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 items-center justify-between">
          <div className="flex items-center gap-8">
            <Link to="/" className="flex items-center gap-2 group">
              <Terminal className="h-6 w-6 text-primary" />
              <div className="text-xl font-bold text-foreground tracking-tight flex items-baseline">
                DoPlan<span className="text-primary font-bold text-base ml-0.5">.dev</span>
              </div>
            </Link>

            <nav className="hidden md:flex items-center gap-6">
              <Link
                to="/commands"
                className="text-sm font-medium text-foreground/80 hover:text-foreground transition-colors"
              >
                Commands
              </Link>
              <Link
                to="/features"
                className="text-sm font-medium text-foreground/80 hover:text-foreground transition-colors"
              >
                Features
              </Link>
            </nav>
          </div>

          <div className="flex items-center gap-3">
            <Button
              variant="ghost"
              size="icon"
              onClick={toggleTheme}
              className="h-9 w-9"
            >
              {actualTheme === 'dark' ? (
                <Sun className="h-4 w-4" />
              ) : (
                <Moon className="h-4 w-4" />
              )}
            </Button>

            <Button
              variant="ghost"
              size="icon"
              asChild
              className="h-9 w-9"
            >
              <a
                href="https://github.com/DoPlan-dev/CLI"
                target="_blank"
                rel="noopener noreferrer"
              >
                <Github className="h-4 w-4" />
              </a>
            </Button>

            <Button
              asChild
              className="bg-primary hover:bg-primary/90 text-primary-foreground gap-2"
            >
              <a href="/docs">
                <BookOpen className="h-4 w-4" />
                Docs
              </a>
            </Button>
          </div>
        </div>
      </div>
    </header>
  );
}
