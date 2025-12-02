import Header from './Header';
import Footer from './Footer';
import { Terminal, Copy, Check, Sparkles, Play, Settings, FileCode, MessageSquare } from 'lucide-react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { useState } from 'react';
import SEO from './SEO';

const workflowCommands = [
    {
        name: 'hey',
        description: 'Start a natural language conversation with your project director. Perfect for brainstorming, asking questions, or getting clarification on project requirements.',
        usage: '/hey [prompt]',
        example: '/hey "How should we structure the database?"',
        icon: MessageSquare,
        color: 'text-blue-500',
        subCommands: [
            { name: 'ask', desc: 'Ask a specific question about the codebase', example: '/hey ask "Where is auth logic?"' },
            { name: 'brainstorm', desc: 'Collaborate on new feature ideas', example: '/hey brainstorm "Viral features"' }
        ]
    },
    {
        name: 'do',
        description: 'Execute a specific task or instruction immediately. The agent will analyze your request, plan the necessary changes, and apply them to your codebase.',
        usage: '/do <instruction>',
        example: '/do "Refactor the authentication middleware"',
        icon: Sparkles,
        color: 'text-cyan-500',
        subCommands: [
            { name: 'fix', desc: 'Identify and fix bugs in a file', example: '/do fix src/App.tsx' },
            { name: 'create', desc: 'Scaffold new components or files', example: '/do create "User profile component"' }
        ]
    },
    {
        name: 'plan',
        description: 'Generate comprehensive implementation plans and architecture. Use this for complex features that require a well-thought-out strategy before coding.',
        usage: '/plan <goal>',
        example: '/plan "Add stripe payment integration"',
        icon: FileCode,
        color: 'text-teal-500',
        subCommands: [
            { name: 'architect', desc: 'Design high-level system architecture', example: '/plan architect "Microservices"' },
            { name: 'roadmap', desc: 'Create a development roadmap', example: '/plan roadmap "Q4 Goals"' }
        ]
    },
    {
        name: 'dev',
        description: 'Enter interactive development mode with live feedback. The agent watches your file changes, runs tests, and suggests fixes in real-time.',
        usage: '/dev',
        example: '/dev',
        icon: Play,
        color: 'text-green-500',
        subCommands: [
            { name: 'watch', desc: 'Monitor files for changes', example: '/dev --watch' },
            { name: 'test', desc: 'Run test suite continuously', example: '/dev --test' }
        ]
    },
    {
        name: 'sys',
        description: 'Manage system operations, deployment, and infrastructure. Handle environment configuration, database migrations, and cloud resources.',
        usage: '/sys <action>',
        example: '/sys deploy production',
        icon: Settings,
        color: 'text-emerald-500',
        subCommands: [
            { name: 'deploy', desc: 'Deploy application to target env', example: '/sys deploy prod' },
            { name: 'logs', desc: 'View system and error logs', example: '/sys logs --tail' },
            { name: 'status', desc: 'Check system health status', example: '/sys status' },
            { name: 'performance', desc: 'View performance metrics and cache statistics', example: '/sys performance' },
            { name: 'backup', desc: 'Create compressed project backups', example: '/sys backup' },
            { name: 'restore', desc: 'Restore from backup', example: '/sys restore' },
            { name: 'memory', desc: 'Export/import memory card', example: '/sys memory' },
            { name: 'engagement', desc: 'View engagement dashboard', example: '/sys engagement' }
        ]
    }
];


interface SubCommand {
    name: string;
    desc: string;
    example: string;
}

interface CommandItem {
    name: string;
    description: string;
    usage: string;
    example: string;
    icon: React.ElementType;
    color: string;
    subCommands?: SubCommand[];
}

function CommandCard({ command, isWorkflow = false }: { command: CommandItem, isWorkflow?: boolean }) {
    const [copied, setCopied] = useState(false);
    const Icon = command.icon || Terminal;

    const copyToClipboard = (text: string) => {
        navigator.clipboard.writeText(text);
        setCopied(true);
        setTimeout(() => setCopied(false), 2000);
    };

    return (
        <Card className={`hover:shadow-lg transition-all duration-300 ${isWorkflow ? 'border-primary/20 bg-primary/5' : ''} w-full`}>
            <CardHeader>
                <div className="flex items-start justify-between">
                    <div className="space-y-1">
                        <CardTitle className="flex items-center gap-3 text-2xl">
                            <div className={`p-2 rounded-lg ${isWorkflow ? 'bg-background shadow-sm' : 'bg-muted'}`}>
                                <Icon className={`h-6 w-6 ${command.color || 'text-primary'}`} />
                            </div>
                            <span>{command.name}</span>
                        </CardTitle>
                        <CardDescription className="text-base max-w-2xl leading-relaxed pt-2">
                            {command.description}
                        </CardDescription>
                    </div>
                </div>
            </CardHeader>
            <CardContent className="space-y-6">
                {/* Main Usage */}
                <div className="relative group">
                    <div className="absolute -top-3 left-3 px-2 bg-background text-xs font-semibold text-muted-foreground">
                        Usage
                    </div>
                    <code className="block bg-background border border-border px-4 py-4 rounded-lg text-sm font-mono pr-12">
                        {command.example}
                    </code>
                    <button
                        onClick={() => copyToClipboard(command.example)}
                        className="absolute right-2 top-1/2 -translate-y-1/2 p-2 rounded hover:bg-muted transition-colors opacity-0 group-hover:opacity-100"
                        aria-label="Copy to clipboard"
                    >
                        {copied ? (
                            <Check className="h-4 w-4 text-green-500" />
                        ) : (
                            <Copy className="h-4 w-4 text-muted-foreground" />
                        )}
                    </button>
                </div>

                {/* Sub Commands */}
                {command.subCommands && (
                    <div className={command.name === 'sys' ? "grid grid-cols-1 md:grid-cols-2 gap-4 pt-2" : "flex flex-col space-y-3 pt-2"}>
                        {command.subCommands.map((sub, idx) => (
                            <div key={idx} className="bg-background/50 border border-border/50 rounded-lg p-3 hover:border-primary/30 transition-colors">
                                <div className="flex items-center justify-between mb-1">
                                    <code className="text-sm font-bold text-primary">{sub.name}</code>
                                </div>
                                <p className="text-xs text-muted-foreground mb-2">{sub.desc}</p>
                                <code className="text-[10px] bg-muted px-1.5 py-0.5 rounded text-muted-foreground block truncate">
                                    {sub.example}
                                </code>
                            </div>
                        ))}
                    </div>
                )}
            </CardContent>
        </Card>
    );
}

export default function CommandsPage() {
    return (
        <div className="min-h-screen bg-background">
            <SEO
                title="Commands"
                description="Master the DoPlan workflow. Comprehensive guide to all CLI commands including /hey, /do, /plan, /dev, and /sys."
                canonical="/commands"
            />
            <Header />
            <main className="container mx-auto px-4 sm:px-6 lg:px-8 py-32">
                <div className="max-w-5xl mx-auto">
                    <div className="text-center mb-16">
                        <h1 className="text-4xl sm:text-5xl font-bold mb-4">
                            Commands & Workflow
                        </h1>
                        <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
                            Master the DoPlan workflow. From natural language interactions to precise system controls.
                        </p>
                    </div>

                    {/* Workflow Commands Section */}
                    <section className="mb-20">
                        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
                            {workflowCommands.map((command, idx) => (
                                <div key={idx} className={idx === workflowCommands.length - 1 ? "lg:col-span-2" : ""}>
                                    <CommandCard command={command} isWorkflow={true} />
                                </div>
                            ))}
                        </div>
                    </section>




                </div>
            </main>
            <Footer />
        </div>
    );
}
