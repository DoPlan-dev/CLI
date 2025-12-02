import Header from './Header';
import Footer from './Footer';
import {
  Zap,
  Rocket,
  Bot,
  BookOpen,
  Paintbrush,
  Plug,
  Brain,
  Trophy,
  Check,
  Award,
  Target
} from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import SEO from './SEO';

const detailedFeatures = [
  {
    icon: Zap,
    title: 'Lightning Fast Setup',
    description: 'Get started in seconds with zero configuration required.',
    benefits: [
      'No installation needed - use npx directly',
      'Automatic dependency detection',
      'Smart project initialization',
      'Instant hot reload during development',
    ],
  },
  {
    icon: Rocket,
    title: 'Production Ready Code',
    description: 'Generate code that follows industry best practices and is ready to deploy.',
    benefits: [
      'TypeScript support out of the box',
      'Optimized build configurations',
      'Security best practices included',
      'Performance optimizations built-in',
    ],
  },
  {
    icon: Bot,
    title: 'Multi-Agent Orchestration',
    description: 'Specialized AI agents collaborate to handle every aspect of development.',
    benefits: [
      'Planning agent for architecture decisions',
      'Development agent for code generation',
      'Testing agent for quality assurance',
      'Deployment agent for CI/CD',
    ],
  },
  {
    icon: BookOpen,
    title: 'Intelligent Documentation',
    description: 'Auto-generated documentation that stays synchronized with your codebase.',
    benefits: [
      'API documentation from code comments',
      'Component usage examples',
      'Automatic changelog generation',
      'Interactive code playground',
    ],
  },
  {
    icon: Paintbrush,
    title: 'Design to Code',
    description: 'Transform your designs into pixel-perfect, production-ready components.',
    benefits: [
      'Figma plugin integration',
      'Sketch file import support',
      'Responsive design generation',
      'Accessibility compliance',
    ],
  },
  {
    icon: Plug,
    title: 'Framework Flexibility',
    description: 'Works seamlessly with your favorite frameworks and tools.',
    benefits: [
      'React, Vue, Angular, Svelte support',
      'Next.js, Nuxt, SvelteKit integration',
      'Tailwind, styled-components, CSS modules',
      'Custom framework adapters',
    ],
  },
  {
    icon: Brain,
    title: 'Context-Aware Intelligence',
    description: 'Understands your entire project to make smart decisions.',
    benefits: [
      'Codebase analysis and understanding',
      'Pattern recognition and replication',
      'Dependency graph awareness',
      'Refactoring suggestions',
    ],
  },
  {
    icon: Trophy,
    title: 'Best Practices Enforcement',
    description: 'Automatically follows coding standards and design patterns.',
    benefits: [
      'ESLint and Prettier integration',
      'SOLID principles adherence',
      'DRY code generation',
      'Clean architecture patterns',
    ],
  },
  {
    icon: Award,
    title: 'Gamified Rewards System',
    description: 'Stay motivated with a comprehensive progression system.',
    benefits: [
      'Level up as you code',
      'Earn badges for achievements',
      'Maintain daily coding streaks',
      'Compete on global leaderboards',
    ],
  },
  {
    icon: Target,
    title: 'Interactive Challenges',
    description: 'Sharpen your skills with AI-generated coding quests.',
    benefits: [
      'Daily coding quests',
      'Skill-based progression paths',
      'Real-world scenario simulations',
      'Instant feedback on solutions',
    ],
  },
];

export default function FeaturesPage() {
  return (
    <div className="min-h-screen bg-background">
      <SEO
        title="Features"
        description="Explore the powerful features of DoPlan: Multi-Agent Orchestration, Intelligent Documentation, Design to Code, and more."
        canonical="/features"
      />
      <Header />

      <main className="pt-24 pb-20">
        {/* Hero Section */}
        <section className="px-4 sm:px-6 lg:px-8 py-16 bg-gradient-to-br from-primary/10 via-transparent to-secondary/10">
          <div className="container mx-auto max-w-4xl text-center">
            <h1 className="text-4xl sm:text-5xl lg:text-6xl font-bold mb-6">
              Powerful Features for{' '}
              <span className="bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
                Modern Development
              </span>
            </h1>
            <p className="text-xl text-muted-foreground">
              Everything you need to build, ship, and scale your applications with AI-powered assistance
            </p>
          </div>
        </section>

        {/* Detailed Features */}
        <section className="px-4 sm:px-6 lg:px-8 py-16">
          <div className="container mx-auto max-w-6xl">
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
              {detailedFeatures.map((feature, index) => {
                const Icon = feature.icon;
                return (
                  <Card
                    key={index}
                    className="border-border/50 hover:shadow-lg transition-all duration-300"
                  >
                    <CardHeader>
                      <div className="flex items-start gap-4">
                        <div className="w-12 h-12 rounded-lg bg-gradient-to-br from-primary/20 to-secondary/20 flex items-center justify-center flex-shrink-0">
                          <Icon className="h-6 w-6 text-primary" />
                        </div>
                        <div>
                          <CardTitle className="text-xl mb-2">{feature.title}</CardTitle>
                          <p className="text-sm text-muted-foreground">
                            {feature.description}
                          </p>
                        </div>
                      </div>
                    </CardHeader>
                    <CardContent>
                      <ul className="space-y-3">
                        {feature.benefits.map((benefit, idx) => (
                          <li key={idx} className="flex items-start gap-3">
                            <Check className="h-5 w-5 text-primary flex-shrink-0 mt-0.5" />
                            <span className="text-sm text-foreground">{benefit}</span>
                          </li>
                        ))}
                      </ul>
                    </CardContent>
                  </Card>
                );
              })}
            </div>
          </div>
        </section>

        {/* Engagement & Growth Section */}
        <section className="px-4 sm:px-6 lg:px-8 py-16 bg-muted/30">
          <div className="container mx-auto max-w-6xl">
            <div className="text-center mb-12">
              <h2 className="text-3xl font-bold mb-4">Engagement & Growth</h2>
              <p className="text-lg text-muted-foreground">
                Gamified progression that keeps you motivated and growing
              </p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-12">
              {/* Rewards */}
              <div className="space-y-6">
                <div className="flex items-center gap-3 mb-6">
                  <Award className="h-8 w-8 text-primary" />
                  <h3 className="text-2xl font-semibold">Rewards System</h3>
                </div>
                <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                  {[
                    { title: 'Badges', desc: 'Visual achievements for milestones' },
                    { title: 'Streaks', desc: 'Track daily coding consistency' },
                    { title: 'Leaderboard', desc: 'Global ranking among developers' },
                    { title: 'XP System', desc: 'Earn experience for every action' }
                  ].map((item, i) => (
                    <div key={i} className="bg-background p-4 rounded-lg border border-border/50">
                      <h4 className="font-semibold text-primary mb-1">{item.title}</h4>
                      <p className="text-sm text-muted-foreground">{item.desc}</p>
                    </div>
                  ))}
                </div>
              </div>

              {/* Challenges */}
              <div className="space-y-6">
                <div className="flex items-center gap-3 mb-6">
                  <Target className="h-8 w-8 text-primary" />
                  <h3 className="text-2xl font-semibold">Interactive Challenges</h3>
                </div>
                <div className="space-y-4">
                  {[
                    { title: 'Daily Quests', desc: 'Quick tasks to keep your skills sharp' },
                    { title: 'Skill Paths', desc: 'Long-term mastery tracks for specific technologies' },
                    { title: 'Bug Bashes', desc: 'Simulated debugging scenarios to test your knowledge' }
                  ].map((item, i) => (
                    <div key={i} className="flex items-start gap-4 bg-background p-4 rounded-lg border border-border/50">
                      <div className="h-2 w-2 rounded-full bg-primary mt-2 shrink-0" />
                      <div>
                        <h4 className="font-semibold mb-1">{item.title}</h4>
                        <p className="text-sm text-muted-foreground">{item.desc}</p>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* CTA Section */}
        <section className="px-4 sm:px-6 lg:px-8 py-16 bg-muted/30">
          <div className="container mx-auto max-w-4xl text-center">
            <h2 className="text-3xl sm:text-4xl font-bold mb-4">
              Explore the Documentation
            </h2>
            <p className="text-lg text-muted-foreground mb-8">
              Dive deep into our guides, API references, and examples to master DoPlan.
            </p>
            <a
              href="/docs"
              className="inline-flex items-center gap-2 px-8 py-4 rounded-lg bg-gradient-to-r from-primary to-secondary text-white font-semibold shadow-lg hover:shadow-xl transition-all duration-300 hover:scale-105 cursor-pointer"
            >
              Read the Docs
            </a>
          </div>
        </section>
      </main>

      <Footer />
    </div>
  );
}
