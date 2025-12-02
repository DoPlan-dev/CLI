import { useState, useEffect } from 'react';
import { X, Cookie } from 'lucide-react';
import { Button } from '@/components/ui/button';

export default function CookieConsent() {
    const [isVisible, setIsVisible] = useState(false);
    const [hasConsented, setHasConsented] = useState(() => {
        return !!localStorage.getItem('cookieConsent');
    });

    useEffect(() => {
        // Show banner after 5 seconds
        const timer = setTimeout(() => {
            setIsVisible(true);
        }, 5000);

        return () => clearTimeout(timer);
    }, []);

    const handleAccept = () => {
        localStorage.setItem('cookieConsent', 'true');
        setHasConsented(true);
        setIsVisible(false);
    };

    const handleDecline = () => {
        localStorage.setItem('cookieConsent', 'false');
        setHasConsented(true);
        setIsVisible(false);
    };

    if (!isVisible || hasConsented) {
        return null;
    }

    return (
        <div className="fixed bottom-0 left-0 right-0 z-50 p-4 sm:p-6 animate-in slide-in-from-bottom duration-500">
            <div className="container mx-auto max-w-6xl">
                <div className="bg-card border border-border rounded-lg shadow-lg p-6">
                    <div className="flex flex-col sm:flex-row items-start sm:items-center gap-4">
                        <div className="flex-shrink-0">
                            <Cookie className="h-8 w-8 text-primary" />
                        </div>

                        <div className="flex-1">
                            <h3 className="text-lg font-semibold mb-2">Cookie Consent</h3>
                            <p className="text-sm text-muted-foreground">
                                We use cookies to enhance your browsing experience, analyze site traffic, and personalize content.
                                By clicking "Accept", you consent to our use of cookies. Read our{' '}
                                <a href="/privacy" className="text-primary hover:underline">
                                    Privacy Policy
                                </a>{' '}
                                for more information.
                            </p>
                        </div>

                        <div className="flex items-center gap-3 w-full sm:w-auto">
                            <Button
                                variant="outline"
                                onClick={handleDecline}
                                className="flex-1 sm:flex-none"
                            >
                                Decline
                            </Button>
                            <Button
                                onClick={handleAccept}
                                className="bg-primary hover:bg-primary/90 text-primary-foreground flex-1 sm:flex-none"
                            >
                                Accept
                            </Button>
                        </div>

                        <button
                            onClick={handleDecline}
                            className="absolute top-4 right-4 text-muted-foreground hover:text-foreground transition-colors"
                            aria-label="Close"
                        >
                            <X className="h-5 w-5" />
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}
