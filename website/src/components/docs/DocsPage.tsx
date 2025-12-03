import { useParams, Navigate } from 'react-router-dom';
import ReactMarkdown from 'react-markdown';
import DocsLayout from './DocsLayout';
import SEO from '../SEO';

// Import all markdown files eagerly to ensure they are bundled
const modules = import.meta.glob('/docs/wiki/**/*.md', { query: '?raw', import: 'default', eager: true });

export default function DocsPage() {
    const { section, page } = useParams();

    // Default redirect to first page if at root /docs
    if (!section && !page) {
        return <Navigate to="/docs/01-Getting-Started/01-Welcome-to-DoPlan" replace />;
    }

    // Construct the path to match the glob pattern
    const path = `/docs/wiki/${section}/${page}.md`;
    const content = modules[path] as string | undefined;

    if (!content) {
        return (
            <DocsLayout>
                <div className="flex flex-col items-center justify-center min-h-[400px] text-center">
                    <h1 className="text-4xl font-bold mb-4">404</h1>
                    <p className="text-xl text-muted-foreground mb-8">Documentation page not found.</p>
                    <p className="text-sm text-muted-foreground">Path: {section}/{page}</p>
                </div>
            </DocsLayout>
        );
    }

    return (
        <DocsLayout>
            <SEO
                title={page ? page.replace(/-/g, ' ') : "Documentation"}
                description="DoPlan Documentation - Learn how to use the AI Project Director."
                canonical={path}
            />
            <div className="animate-in fade-in duration-500">
                <ReactMarkdown
                    components={{
                        h1: ({ ...props }) => <h1 className="text-3xl font-bold mb-6 pb-2 border-b border-border" {...props} />,
                        h2: ({ ...props }) => <h2 className="text-2xl font-semibold mt-8 mb-4" {...props} />,
                        h3: ({ ...props }) => <h3 className="text-xl font-semibold mt-6 mb-3" {...props} />,
                        p: ({ ...props }) => <p className="leading-7 mb-4 text-muted-foreground" {...props} />,
                        ul: ({ ...props }) => <ul className="list-disc list-outside ml-6 mb-4 space-y-1 text-muted-foreground" {...props} />,
                        ol: ({ ...props }) => <ol className="list-decimal list-outside ml-6 mb-4 space-y-1 text-muted-foreground" {...props} />,
                        li: ({ ...props }) => <li className="pl-1" {...props} />,
                        a: ({ ...props }) => <a className="text-primary hover:underline font-medium" {...props} />,
                        code: ({ inline, className, children, ...props }: React.ComponentPropsWithoutRef<'code'> & { inline?: boolean }) => {
                            const match = /language-(\w+)/.exec(className || '');
                            return !inline ? (
                                <div className="relative group my-4">
                                    <div className="absolute -top-3 left-3 px-2 bg-muted text-xs font-semibold text-muted-foreground rounded border border-border">
                                        {match ? match[1] : 'code'}
                                    </div>
                                    <pre className="bg-muted/50 border border-border rounded-lg p-4 overflow-x-auto">
                                        <code className={className} {...props}>
                                            {children}
                                        </code>
                                    </pre>
                                </div>
                            ) : (
                                <code className="bg-muted px-1.5 py-0.5 rounded text-sm font-mono text-foreground" {...props}>
                                    {children}
                                </code>
                            );
                        },
                        blockquote: ({ ...props }) => (
                            <blockquote className="border-l-4 border-primary/30 pl-4 italic my-4 text-muted-foreground bg-muted/20 py-2 pr-4 rounded-r" {...props} />
                        ),
                        table: ({ ...props }) => (
                            <div className="overflow-x-auto my-6 border border-border rounded-lg">
                                <table className="w-full text-sm text-left" {...props} />
                            </div>
                        ),
                        thead: ({ ...props }) => <thead className="bg-muted/50 text-xs uppercase text-muted-foreground" {...props} />,
                        th: ({ ...props }) => <th className="px-6 py-3 font-medium" {...props} />,
                        td: ({ ...props }) => <td className="px-6 py-4 border-t border-border" {...props} />,
                    }}
                >
                    {content}
                </ReactMarkdown>
            </div>
        </DocsLayout>
    );
}
