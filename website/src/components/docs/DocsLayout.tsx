import { useState, useEffect } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Menu, X, ChevronRight, Github, Moon, Sun, Laptop, Book } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import Header from '../Header';
import Footer from '../Footer';

// Navigation structure based on wiki INDEX.md
const navigation = [
    {
        title: 'Getting Started',
        path: '01-Getting-Started',
        items: [
            { title: 'Welcome to DoPlan', path: '01-Welcome-to-DoPlan' },
            { title: 'Installation', path: '02-Installation' },
            { title: 'First Project', path: '03-First-Project' },
            { title: 'Quick Tour', path: '04-Quick-Tour' },
        ]
    },
    {
        title: 'Commands',
        path: '02-Commands',
        items: [
            { title: 'Command Overview', path: '01-Command-Overview' },
            { title: 'Workflow Commands', path: '02-Workflow-Commands' },
            { title: 'System Commands', path: '03-System-Commands' },
            { title: 'Command Reference', path: '04-Command-Reference' },
        ]
    },
    {
        title: 'Engagement System',
        path: '03-Engagement-System',
        items: [
            { title: 'Overview', path: '01-Overview' },
            { title: 'Achievements', path: '02-Achievements' },
            { title: 'Challenges', path: '03-Challenges' },
            { title: 'Score System', path: '04-Score-System' },
            { title: 'Reward System', path: '05-Reward-System' },
            { title: 'Engagement Dashboard', path: '06-Engagement-Dashboard' },
        ]
    },
    {
        title: 'Memory and Brain',
        path: '04-Memory-and-Brain',
        items: [
            { title: 'Overview', path: '01-Overview' },
            { title: 'Memory Card', path: '02-Memory-Card' },
            { title: 'Brain System', path: '03-Brain-System' },
            { title: 'Personalization', path: '04-Personalization' },
            { title: 'Relationship Building', path: '05-Relationship-Building' },
        ]
    },
    {
        title: 'Workflow',
        path: '05-Workflow',
        items: [
            { title: 'Complete Workflow', path: '01-Complete-Workflow' },
            { title: 'Phase-by-Phase Guide', path: '02-Phase-by-Phase' },
            { title: 'Best Practices', path: '03-Best-Practices' },
            { title: 'Common Patterns', path: '04-Common-Patterns' },
        ]
    },
    {
        title: 'Features',
        path: '06-Features',
        items: [
            { title: 'Time Tracking', path: '01-Time-Tracking' },
            { title: 'State Management', path: '02-State-Management' },
            { title: 'Git Automation', path: '03-Git-Automation' },
            { title: 'Learning Support', path: '04-Learning-Support' },
            { title: 'Personalization', path: '05-Personalization' },
            { title: 'Backup and Restore', path: '06-Backup-and-Restore' },
        ]
    },
    {
        title: 'Learning & Education',
        path: '07-Learning-Education',
        items: [
            { title: 'For Beginners', path: '01-For-Beginners' },
            { title: 'For Intermediate', path: '02-For-Intermediate' },
            { title: 'For Advanced', path: '03-For-Advanced' },
            { title: 'Learning Goals', path: '04-Learning-Goals' },
            { title: 'Tech Exploration', path: '05-Tech-Exploration' },
        ]
    },
    {
        title: 'Advanced Topics',
        path: '08-Advanced',
        items: [
            { title: 'System Control', path: '01-System-Control' },
            { title: 'Customization', path: '02-Customization' },
            { title: 'Integration', path: '03-Integration' },
            { title: 'Troubleshooting', path: '04-Troubleshooting' },
            { title: 'Contributing', path: '05-Contributing' },
        ]
    },
    {
        title: 'Reference',
        path: '09-Reference',
        items: [
            { title: 'Command Quick Reference', path: '01-Command-Quick-Reference' },
            { title: 'File Structure', path: '02-File-Structure' },
            { title: 'State Management', path: '03-State-Management' },
            { title: 'Time Tracker Format', path: '04-Time-Tracker-Format' },
            { title: 'Memory Card Schema', path: '05-Memory-Card-Schema' },
        ]
    },
];

export default function DocsLayout({ children }: { children: React.ReactNode }) {
    const [isSidebarOpen, setIsSidebarOpen] = useState(false);
    const location = useLocation();

    // Close sidebar on route change (mobile)
    useEffect(() => {
        if (isSidebarOpen) {
            setIsSidebarOpen(false);
        }
    }, [location.pathname]);

    const currentPath = location.pathname.replace('/docs/', '');

    return (
        <div className="min-h-screen bg-background flex flex-col">
            <Header />

            <div className="flex-1 container mx-auto px-4 sm:px-6 lg:px-8 pt-24 pb-12 flex items-start gap-8">
                {/* Mobile Sidebar Toggle */}
                <Button
                    variant="outline"
                    size="icon"
                    className="lg:hidden fixed bottom-6 right-6 z-50 rounded-full shadow-lg h-12 w-12"
                    onClick={() => setIsSidebarOpen(!isSidebarOpen)}
                >
                    {isSidebarOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
                </Button>

                {/* Sidebar */}
                <aside
                    className={`
            fixed inset-y-0 left-0 z-40 w-64 bg-background border-r border-border transform transition-transform duration-300 ease-in-out lg:translate-x-0 lg:static lg:h-[calc(100vh-8rem)] lg:border-none lg:bg-transparent
            ${isSidebarOpen ? 'translate-x-0' : '-translate-x-full'}
          `}
                >
                    <div className="h-full flex flex-col pt-20 lg:pt-0">
                        <div className="px-4 mb-4 lg:hidden">
                            <div className="flex items-center gap-2 font-bold text-xl">
                                <Book className="h-6 w-6 text-primary" />
                                Documentation
                            </div>
                        </div>

                        <ScrollArea className="flex-1 px-2">
                            <div className="space-y-6 pb-10">
                                {navigation.map((section, idx) => (
                                    <div key={idx}>
                                        <h4 className="px-2 mb-2 text-sm font-semibold text-foreground flex items-center gap-2">
                                            {section.title}
                                        </h4>
                                        <div className="space-y-1">
                                            {section.items.map((item, itemIdx) => {
                                                const itemPath = `${section.path}/${item.path}`;
                                                const isActive = currentPath.includes(itemPath);

                                                return (
                                                    <Link
                                                        key={itemIdx}
                                                        to={`/docs/${itemPath}`}
                                                        className={`
                              group flex items-center justify-between px-2 py-1.5 text-sm rounded-md transition-colors
                              ${isActive
                                                                ? 'bg-primary/10 text-primary font-medium'
                                                                : 'text-muted-foreground hover:bg-muted hover:text-foreground'
                                                            }
                            `}
                                                    >
                                                        {item.title}
                                                        {isActive && <ChevronRight className="h-3 w-3" />}
                                                    </Link>
                                                );
                                            })}
                                        </div>
                                    </div>
                                ))}
                            </div>
                        </ScrollArea>
                    </div>
                </aside>

                {/* Overlay for mobile sidebar */}
                {isSidebarOpen && (
                    <div
                        className="fixed inset-0 bg-background/80 backdrop-blur-sm z-30 lg:hidden"
                        onClick={() => setIsSidebarOpen(false)}
                    />
                )}

                {/* Main Content */}
                <main className="flex-1 min-w-0 w-full">
                    <div className="prose prose-slate dark:prose-invert max-w-none">
                        {children}
                    </div>
                </main>
            </div>

            <Footer />
        </div>
    );
}
