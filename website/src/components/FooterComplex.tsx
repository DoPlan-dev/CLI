import { Github, Twitter, Linkedin } from 'lucide-react';
import { Link } from 'react-router-dom';

export default function Footer() {
    return (
        <footer className="border-t border-border/40 bg-muted/30">
            <div className="container mx-auto px-4 sm:px-6 lg:px-8 py-12">
                <div className="grid grid-cols-1 md:grid-cols-4 gap-8 mb-8">
                    {/* Brand */}
                    <div className="space-y-4">
                        <div className="text-2xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
                            DoPlan.dev
                        </div>
                        <p className="text-sm text-muted-foreground">
                            The AI-powered project director that transforms ideas into production-ready code.
                        </p>
                        <div className="flex items-center gap-3">
                            <a
                                href="https://github.com/DoPlan-dev/CLI"
                                target="_blank"
                                rel="noopener noreferrer"
                                className="text-muted-foreground hover:text-foreground transition-colors"
                            >
                                <Github className="h-5 w-5" />
                            </a>
                            <a
                                href="https://twitter.com/doplan"
                                target="_blank"
                                rel="noopener noreferrer"
                                className="text-muted-foreground hover:text-foreground transition-colors"
                            >
                                <Twitter className="h-5 w-5" />
                            </a>
                            <a
                                href="https://linkedin.com/company/doplan"
                                target="_blank"
                                rel="noopener noreferrer"
                                className="text-muted-foreground hover:text-foreground transition-colors"
                            >
                                <Linkedin className="h-5 w-5" />
                            </a>
                        </div>
                    </div>

                    {/* Product */}
                    <div>
                        <h3 className="font-semibold mb-4 text-foreground">Product</h3>
                        <ul className="space-y-2">
                            <li>
                                <Link to="/features" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Features
                                </Link>
                            </li>
                            <li>
                                <a href="/docs" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Documentation
                                </a>
                            </li>
                            <li>
                                <a href="/docs/examples" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Examples
                                </a>
                            </li>
                            <li>
                                <a href="/docs/reference" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    API Reference
                                </a>
                            </li>
                        </ul>
                    </div>

                    {/* Resources */}
                    <div>
                        <h3 className="font-semibold mb-4 text-foreground">Resources</h3>
                        <ul className="space-y-2">
                            <li>
                                <a href="/docs/getting-started" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Getting Started
                                </a>
                            </li>
                            <li>
                                <a href="/docs/guides" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Guides
                                </a>
                            </li>
                            <li>
                                <a href="https://github.com/doplan-adv/cli" target="_blank" rel="noopener noreferrer" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    GitHub
                                </a>
                            </li>
                            <li>
                                <a href="/docs/troubleshooting" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Troubleshooting
                                </a>
                            </li>
                        </ul>
                    </div>

                    {/* Company */}
                    <div>
                        <h3 className="font-semibold mb-4 text-foreground">Company</h3>
                        <ul className="space-y-2">
                            <li>
                                <a href="/about" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    About
                                </a>
                            </li>
                            <li>
                                <a href="/blog" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Blog
                                </a>
                            </li>
                            <li>
                                <a href="/careers" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Careers
                                </a>
                            </li>
                            <li>
                                <a href="/contact" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                                    Contact
                                </a>
                            </li>
                        </ul>
                    </div>
                </div>

                {/* Bottom Bar */}
                <div className="pt-8 border-t border-border/40 flex flex-col sm:flex-row justify-between items-center gap-4">
                    <p className="text-sm text-muted-foreground">
                        © {new Date().getFullYear()} DoPlan.dev. All rights reserved.
                    </p>
                    <div className="flex items-center gap-6">
                        <a href="/privacy" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                            Privacy Policy
                        </a>
                        <a href="/terms" className="text-sm text-muted-foreground hover:text-foreground transition-colors">
                            Terms of Service
                        </a>
                    </div>
                </div>
            </div>
        </footer>
    );
}
