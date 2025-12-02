import { ArrowRight, Hand, Lightbulb, FileCheck, Wrench, Settings, ChevronRight } from 'lucide-react';
import { motion } from 'framer-motion';

const steps = [
  {
    name: 'Hey',
    description: 'Describe what you want to build',
    color: 'from-blue-500 to-cyan-500',
    icon: Hand,
  },
  {
    name: 'Do',
    description: 'AI agents analyze requirements',
    color: 'from-cyan-500 to-teal-500',
    icon: Lightbulb,
  },
  {
    name: 'Plan',
    description: 'Generate architecture & tasks',
    color: 'from-teal-500 to-green-500',
    icon: FileCheck,
  },
  {
    name: 'Dev',
    description: 'Write production-ready code',
    color: 'from-green-500 to-emerald-500',
    icon: Wrench,
  },
  {
    name: 'Sys',
    description: 'Test, deploy, and monitor',
    color: 'from-emerald-500 to-lime-500',
    icon: Settings,
  },
];

const containerVariants = {
  hidden: { opacity: 0 },
  visible: {
    opacity: 1,
    transition: {
      staggerChildren: 0.3
    }
  }
};

export default function Workflow() {
  return (
    <section className="py-20 px-4 sm:px-6 lg:px-8 min-h-[1080px] flex items-center">
      <div className="container mx-auto">
        <motion.div
          className="max-w-3xl mx-auto text-center mb-16"
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true, margin: "-100px" }}
          transition={{ duration: 0.8, ease: "easeOut" }}
        >
          <h2 className="text-3xl sm:text-4xl lg:text-5xl font-bold mb-4">
            Your AI-Powered{' '}
            <span className="bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
              Development Pipeline
            </span>
          </h2>
          <p className="text-lg text-muted-foreground">
            From idea to production in five intelligent steps
          </p>
        </motion.div>

        <motion.div
          className="max-w-6xl mx-auto"
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, margin: "-100px" }}
          variants={containerVariants}
        >
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4 md:gap-2">
            {steps.map((step, index) => {
              const Icon = step.icon;
              return (
                <div key={index} className="flex flex-col items-center">
                  <div className="relative group">
                    {/* Step Card */}
                    <div className="relative z-10 p-6 rounded-2xl bg-card border border-border shadow-lg hover:shadow-xl transition-all duration-300 hover:-translate-y-2">
                      <div className={`w-16 h-16 mx-auto mb-4 rounded-xl bg-gradient-to-br ${step.color} flex items-center justify-center shadow-lg`}>
                        <Icon className="h-8 w-8 text-white" />
                      </div>
                      <h3 className="text-xl font-bold text-center mb-2 text-foreground">
                        {step.name}
                      </h3>
                      <p className="text-sm text-center text-muted-foreground leading-relaxed">
                        {step.description}
                      </p>
                    </div>

                    {/* Arrow (hidden on last item and mobile) */}
                    {index < steps.length - 1 && (
                      <div className="hidden md:block absolute top-1/2 -right-6 transform -translate-y-1/2 z-20">
                        <ChevronRight className="h-12 w-12 text-primary/50" strokeWidth={4} />
                      </div>
                    )}
                  </div>
                </div>
              );
            })}
          </div>
        </motion.div>

        {/* Bottom CTA */}
        <div className="mt-16 text-center">
          <p className="text-muted-foreground mb-4">
            Stay up to date with our latest improvements
          </p>
          <a
            href="https://github.com/DoPlan-dev/CLI/releases/"
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center gap-2 px-6 py-3 rounded-lg bg-gradient-to-r from-primary to-secondary text-white font-semibold shadow-lg hover:shadow-xl transition-all duration-300 hover:scale-105 cursor-pointer"
          >
            View Release Notes
            <ArrowRight className="h-4 w-4" />
          </a>
        </div>
      </div>
    </section >
  );
}
