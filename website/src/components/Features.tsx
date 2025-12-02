import {
  Zap,
  Rocket,
  Bot,
  BookOpen,
  Paintbrush,
  Plug,
  Brain,
  Trophy,
  Cloud,
  Users,
  Shield,
  Smartphone,
  Globe,
  Code
} from 'lucide-react';
import { Card, CardContent } from '@/components/ui/card';
import { motion } from 'framer-motion';
import { useRef, useState, useEffect } from 'react';

const features = [
  {
    icon: Zap,
    title: 'Lightning Fast',
    description: 'Zero-install setup gets you coding in seconds. No configuration files, no boilerplate hell.',
  },
  {
    icon: Rocket,
    title: 'Production Ready',
    description: 'Generate production-grade code with best practices baked in. Deploy with confidence.',
  },
  {
    icon: Bot,
    title: 'Multi-Agent System',
    description: 'Specialized AI agents work together—planning, coding, testing, and deploying your project.',
  },
  {
    icon: BookOpen,
    title: 'Smart Documentation',
    description: 'Auto-generated docs that stay in sync with your code. Always up-to-date, always accurate.',
  },
  {
    icon: Paintbrush,
    title: 'Design to Code',
    description: 'Transform designs into pixel-perfect components. Support for Figma, Sketch, and more.',
  },
  {
    icon: Plug,
    title: 'Framework Agnostic',
    description: 'Works with React, Vue, Next.js, and more. Use your favorite tools and frameworks.',
  },
  {
    icon: Brain,
    title: 'Context Aware',
    description: 'Understands your entire codebase. Makes intelligent decisions based on your project context.',
  },
  {
    icon: Trophy,
    title: 'Best Practices',
    description: 'Follows industry standards and patterns. Clean, maintainable, and scalable code every time.',
  },
  {
    icon: Cloud,
    title: 'Cloud Sync',
    description: 'Seamlessly sync your projects across devices. Work from anywhere, anytime.',
  },
  {
    icon: Users,
    title: 'Team Collaboration',
    description: 'Built for teams. Share agents, workflows, and project context with your colleagues.',
  },
  {
    icon: Shield,
    title: 'Enterprise Security',
    description: 'Bank-grade security for your code and data. SOC2 compliant and ready for enterprise.',
  },
  {
    icon: Smartphone,
    title: 'Mobile First',
    description: 'Generate responsive, mobile-first applications by default. Look great on any device.',
  },
  {
    icon: Globe,
    title: 'Global CDN',
    description: 'Deploy to a global edge network for lightning-fast performance worldwide.',
  },
  {
    icon: Code,
    title: 'Custom Agents',
    description: 'Create your own specialized agents to automate specific tasks in your workflow.',
  },
];

const FeatureCard = ({ feature }: { feature: typeof features[0] }) => {
  const Icon = feature.icon;
  return (
    <Card className="w-[350px] shrink-0 group hover:shadow-lg transition-all duration-300 border-border/50 bg-card/50 backdrop-blur mx-4 hover:bg-[#7149e4]">
      <CardContent className="p-6 space-y-4">
        <div className="w-12 h-12 rounded-lg bg-gradient-to-br from-primary/20 to-secondary/20 flex items-center justify-center group-hover:scale-110 transition-transform group-hover:bg-white/10">
          <Icon className="h-6 w-6 text-primary group-hover:text-white" />
        </div>
        <h3 className="text-lg font-semibold text-foreground group-hover:text-white">
          {feature.title}
        </h3>
        <p className="text-sm text-muted-foreground leading-relaxed group-hover:text-white/80">
          {feature.description}
        </p>
      </CardContent>
    </Card>
  );
};

export default function Features() {
  const containerRef = useRef<HTMLDivElement>(null);
  const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });
  const [circuitPaths, setCircuitPaths] = useState<any[]>([]);

  useEffect(() => {
    // Generate random circuit paths on client side only
    const paths = [...Array(20)].map(() => {
      const startX = Math.random() * 100;
      const startY = Math.random() * 100;
      const segments = Math.floor(Math.random() * 3) + 2;

      let d = `M${startX}%,${startY}%`;
      let currentX = startX;
      let currentY = startY;

      for (let j = 0; j < segments; j++) {
        const isHorizontal = Math.random() > 0.5;
        const length = (Math.random() * 20 + 10) * (Math.random() > 0.5 ? 1 : -1);

        if (isHorizontal) {
          currentX += length;
          d += ` H${currentX}%`;
        } else {
          currentY += length;
          d += ` V${currentY}%`;
        }
      }

      return {
        d,
        duration: Math.random() * 5 + 3,
        delay: Math.random() * 5,
        repeatDelay: Math.random() * 3
      };
    });
    setCircuitPaths(paths);
  }, []);

  const handleMouseMove = (e: React.MouseEvent<HTMLDivElement>) => {
    if (containerRef.current) {
      const rect = containerRef.current.getBoundingClientRect();
      setMousePosition({
        x: e.clientX - rect.left,
        y: e.clientY - rect.top,
      });
    }
  };

  const firstRow = features.slice(0, 7);
  const secondRow = features.slice(7);

  return (
    <section
      ref={containerRef}
      onMouseMove={handleMouseMove}
      className="py-20 bg-background min-h-[1080px] flex flex-col justify-center overflow-hidden relative"
    >
      {/* Interactive Spotlight Grid */}
      <div className="absolute inset-0 pointer-events-none">
        {/* Base Grid */}
        <div
          className="absolute inset-0 opacity-[0.08]"
          style={{
            backgroundImage: `linear-gradient(to right, #888 1px, transparent 1px), linear-gradient(to bottom, #888 1px, transparent 1px)`,
            backgroundSize: '40px 40px'
          }}
        />

        {/* Spotlight Reveal */}
        <div
          className="absolute inset-0 opacity-10 transition-opacity duration-300"
          style={{
            background: `radial-gradient(600px circle at ${mousePosition.x}px ${mousePosition.y}px, rgba(var(--primary-rgb), 0.15), transparent 40%)`,
          }}
        />
      </div>

      {/* Circuit Board Traces */}
      <div className="absolute inset-0 pointer-events-none opacity-20">
        <svg className="w-full h-full" preserveAspectRatio="none">
          <defs>
            <linearGradient id="trace-gradient" x1="0%" y1="0%" x2="100%" y2="0%">
              <stop offset="0%" stopColor="transparent" />
              <stop offset="50%" stopColor="hsl(var(--primary))" />
              <stop offset="100%" stopColor="transparent" />
            </linearGradient>
          </defs>

          {circuitPaths.map((path, i) => (
            <motion.path
              key={i}
              d={path.d}
              fill="none"
              stroke="url(#trace-gradient)"
              strokeWidth="2"
              initial={{ pathLength: 0, opacity: 0 }}
              animate={{
                pathLength: [0, 1, 1],
                pathOffset: [0, 0, 1],
                opacity: [0, 1, 0]
              }}
              transition={{
                duration: path.duration,
                repeat: Infinity,
                ease: "linear",
                repeatDelay: path.repeatDelay,
                delay: path.delay
              }}
            />
          ))}
        </svg>
      </div>

      <div className="container mx-auto px-4 sm:px-6 lg:px-8 mb-16 relative z-10">
        <div className="max-w-3xl mx-auto text-center">
          <h2 className="text-3xl sm:text-4xl lg:text-5xl font-bold mb-4">
            <span className="bg-gradient-to-r from-primary via-secondary to-primary bg-clip-text text-transparent">
              Everything You Need
            </span>
            <br />
            to Ship Faster
          </h2>
          <p className="text-lg text-muted-foreground">
            Powerful features that streamline your development workflow from start to finish
          </p>
        </div>
      </div>

      <div className="space-y-8">
        {/* Top Row - Moves Right */}
        <div className="relative flex overflow-hidden">
          <motion.div
            className="flex"
            animate={{ x: ["-33.33%", "0%"] }}
            transition={{
              x: {
                repeat: Infinity,
                repeatType: "loop",
                duration: 120,
                ease: "linear",
              },
            }}
          >
            {[...firstRow, ...firstRow, ...firstRow].map((feature, index) => (
              <FeatureCard key={`row1-${index}`} feature={feature} />
            ))}
          </motion.div>
        </div>

        {/* Bottom Row - Moves Left */}
        <div className="relative flex overflow-hidden">
          <motion.div
            className="flex"
            animate={{ x: ["0%", "-33.33%"] }}
            transition={{
              x: {
                repeat: Infinity,
                repeatType: "loop",
                duration: 120,
                ease: "linear",
              },
            }}
          >
            {[...secondRow, ...secondRow, ...secondRow].map((feature, index) => (
              <FeatureCard key={`row2-${index}`} feature={feature} />
            ))}
          </motion.div>
        </div>
      </div>
    </section>
  );
}
