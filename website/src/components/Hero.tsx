import Terminal from './Terminal';
import AnimatedBackground from './AnimatedBackground';
import { ArrowRight, Sparkles } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { motion } from 'framer-motion';

const fadeInUp = {
  initial: { opacity: 0, y: 30 },
  animate: { opacity: 1, y: 0 },
  transition: { duration: 1.0, ease: "easeOut" }
};

const staggerContainer = {
  animate: {
    transition: {
      staggerChildren: 0.3
    }
  }
};

export default function Hero() {
  return (
    <section className="relative pt-32 pb-20 px-4 sm:px-6 lg:px-8 overflow-hidden min-h-[1080px] flex items-center bg-background">
      {/* Animated particle background */}
      <AnimatedBackground />

      <motion.div
        className="container mx-auto relative"
        initial="initial"
        animate="animate"
        variants={staggerContainer}
      >
        <div className="max-w-4xl mx-auto text-center space-y-8">
          {/* Badge */}
          <motion.div variants={fadeInUp} className="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-primary/10 border border-primary/20 text-sm font-medium text-primary">
            <Sparkles className="h-4 w-4" />
            Zero-install AI Project Director
          </motion.div>

          {/* Heading */}
          <motion.h1 variants={fadeInUp} className="text-4xl sm:text-5xl lg:text-6xl font-bold tracking-tight">
            <span className="bg-gradient-to-r from-primary via-secondary to-primary bg-clip-text text-transparent">
              Build Faster
            </span>
            <br />
            <span className="text-foreground">
              With AI Agents
            </span>
          </motion.h1>

          {/* Description */}
          <motion.p variants={fadeInUp} className="text-lg text-muted-foreground max-w-2xl mx-auto leading-relaxed">
            DoPlan orchestrates multiple AI agents to handle your entire development workflow—from planning to deployment. No configuration, no setup, just results.
          </motion.p>

          {/* Terminal */}
          <motion.div variants={fadeInUp} className="pt-8">
            <Terminal />
          </motion.div>

          {/* CTA Buttons */}
          <div className="flex flex-col sm:flex-row items-center justify-center gap-4 pt-4">
            <Button
              size="lg"
              className="bg-primary hover:bg-primary/90 text-primary-foreground group"
              asChild
            >
              <a href="/docs">
                Documentation
                <ArrowRight className="ml-2 h-4 w-4 group-hover:translate-x-1 transition-transform" />
              </a>
            </Button>
            <Button
              size="lg"
              variant="outline"
              asChild
            >
              <a href="https://github.com/DoPlan-dev/CLI" target="_blank" rel="noopener noreferrer">
                View on GitHub
              </a>
            </Button>
          </div>

          {/* Stats */}
          <div className="flex justify-center pt-12 max-w-2xl mx-auto">
            <div className="space-y-1">
              <div className="text-3xl font-bold text-foreground">v1.3.4</div>
              <div className="text-sm text-muted-foreground">Latest</div>
            </div>
          </div>
        </div>
      </motion.div>
    </section>
  );
}
