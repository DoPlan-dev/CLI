import { useEffect, useState } from 'react';
import Header from './Header';
import Footer from './Footer';
import { cn } from '@/lib/utils';
import SEO from './SEO';

const sections = [
    { id: 'acceptance', title: '1. Acceptance of Terms' },
    { id: 'description', title: '2. Description of Service' },
    { id: 'accounts', title: '3. User Accounts' },
    { id: 'subscriptions', title: '4. Subscriptions & Billing' },
    { id: 'acceptable-use', title: '5. Acceptable Use' },
    { id: 'intellectual-property', title: '6. Intellectual Property' },
    { id: 'user-content', title: '7. User Content & Code' },
    { id: 'third-party', title: '8. Third-Party Links' },
    { id: 'termination', title: '9. Termination' },
    { id: 'indemnification', title: '10. Indemnification' },
    { id: 'limitation', title: '11. Limitation of Liability' },
    { id: 'disclaimer', title: '12. Disclaimer' },
    { id: 'governing-law', title: '13. Governing Law' },
    { id: 'changes', title: '14. Changes to Terms' },
    { id: 'contact', title: '15. Contact Us' },
];

export default function TermsOfService() {
    const [activeSection, setActiveSection] = useState('acceptance');

    useEffect(() => {
        const handleScroll = () => {
            const sectionElements = sections.map(section => document.getElementById(section.id));
            const scrollPosition = window.scrollY + 200;

            for (let i = sectionElements.length - 1; i >= 0; i--) {
                const section = sectionElements[i];
                if (section && section.offsetTop <= scrollPosition) {
                    setActiveSection(sections[i].id);
                    break;
                }
            }
        };

        window.addEventListener('scroll', handleScroll);
        return () => window.removeEventListener('scroll', handleScroll);
    }, []);

    const scrollToSection = (id: string) => {
        const element = document.getElementById(id);
        if (element) {
            window.scrollTo({
                top: element.offsetTop - 100,
                behavior: 'smooth'
            });
        }
    };

    return (
        <div className="min-h-screen bg-background">
            <SEO
                title="Terms of Service"
                description="DoPlan Terms of Service - Rules and regulations for using our services."
                canonical="/terms"
            />
            <Header />
            <main className="container mx-auto px-4 sm:px-6 lg:px-8 py-32">
                <div className="flex flex-col lg:flex-row gap-12 max-w-7xl mx-auto">
                    {/* Sidebar Navigation */}
                    <aside className="lg:w-1/4 hidden lg:block">
                        <div className="sticky top-32 space-y-1">
                            <h3 className="font-semibold mb-4 px-4 text-lg">Table of Contents</h3>
                            {sections.map((section) => (
                                <button
                                    key={section.id}
                                    onClick={() => scrollToSection(section.id)}
                                    className={cn(
                                        "block w-full text-left px-4 py-2 text-sm rounded-md transition-colors",
                                        activeSection === section.id
                                            ? "bg-primary/10 text-primary font-medium"
                                            : "text-muted-foreground hover:bg-muted hover:text-foreground"
                                    )}
                                >
                                    {section.title}
                                </button>
                            ))}
                        </div>
                    </aside>

                    {/* Main Content */}
                    <div className="lg:w-3/4">
                        <h1 className="text-4xl sm:text-5xl font-bold mb-6">Terms of Service</h1>
                        <p className="text-muted-foreground mb-12 text-lg">
                            Last updated: {new Date().toLocaleDateString()}
                        </p>

                        <div className="prose prose-lg dark:prose-invert max-w-none space-y-16">
                            <section id="acceptance" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">1. Acceptance of Terms</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    By accessing or using the DoPlan.dev website, CLI tools, API, and related services (collectively, the "Service"), you agree to be bound by these Terms of Service ("Terms"). These Terms apply to all visitors, users, developers, and others who access or use the Service.
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    If you disagree with any part of the terms, then you may not access the Service. Your access to and use of the Service is conditioned on your acceptance of and compliance with these Terms.
                                </p>
                            </section>

                            <section id="description" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">2. Description of Service</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    DoPlan.dev provides an AI-powered development pipeline that orchestrates multiple AI agents to handle your entire development workflow—from planning to deployment. The Service includes but is not limited to:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li>AI-driven code generation and refactoring</li>
                                    <li>Project planning and architecture design</li>
                                    <li>Automated testing and deployment pipelines</li>
                                    <li>CLI tools for local development integration</li>
                                    <li>Web dashboard for project management</li>
                                </ul>
                            </section>

                            <section id="accounts" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">3. User Accounts</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    When you create an account with us, you must provide us information that is accurate, complete, and current at all times. Failure to do so constitutes a breach of the Terms, which may result in immediate termination of your account on our Service.
                                </p>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    You are responsible for safeguarding the password that you use to access the Service and for any activities or actions under your password, whether your password is with our Service or a third-party service. You agree not to disclose your password to any third party.
                                </p>
                            </section>

                            <section id="subscriptions" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">4. Subscriptions and Billing</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    Some parts of the Service are billed on a subscription basis ("Subscription(s)"). You will be billed in advance on a recurring and periodic basis ("Billing Cycle"). Billing cycles are set either on a monthly or annual basis, depending on the type of subscription plan you select when purchasing a Subscription.
                                </p>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    At the end of each Billing Cycle, your Subscription will automatically renew under the exact same conditions unless you cancel it or DoPlan.dev cancels it. You may cancel your Subscription renewal either through your online account management page or by contacting DoPlan.dev customer support team.
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    A valid payment method, including credit card, is required to process the payment for your Subscription. You shall provide DoPlan.dev with accurate and complete billing information including full name, address, state, zip code, telephone number, and a valid payment method information.
                                </p>
                            </section>

                            <section id="acceptable-use" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">5. Acceptable Use</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    You agree not to use the Service:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li>In any way that violates any applicable national or international law or regulation.</li>
                                    <li>For the purpose of exploiting, harming, or attempting to exploit or harm minors in any way.</li>
                                    <li>To transmit, or procure the sending of, any advertising or promotional material, including any "junk mail", "chain letter," "spam," or any other similar solicitation.</li>
                                    <li>To impersonate or attempt to impersonate DoPlan.dev, a DoPlan.dev employee, another user, or any other person or entity.</li>
                                    <li>To engage in any other conduct that restricts or inhibits anyone's use or enjoyment of the Service, or which, as determined by us, may harm DoPlan.dev or users of the Service or expose them to liability.</li>
                                </ul>
                            </section>

                            <section id="intellectual-property" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">6. Intellectual Property</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    The Service and its original content (excluding Content provided by users), features and functionality are and will remain the exclusive property of DoPlan.dev and its licensors. The Service is protected by copyright, trademark, and other laws of both the United States and foreign countries.
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    Our trademarks and trade dress may not be used in connection with any product or service without the prior written consent of DoPlan.dev.
                                </p>
                            </section>

                            <section id="user-content" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">7. User Content & Code Ownership</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    You retain full ownership of all code, assets, projects, and intellectual property generated using DoPlan.dev ("User Content"). We claim no intellectual property rights over the material you provide to the Service or the code generated by the Service for your specific projects.
                                </p>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    By using the Service, you grant DoPlan.dev a worldwide, non-exclusive, royalty-free license to use, reproduce, modify, adapt, publish, translate, and distribute your User Content solely for the purpose of providing and improving the Service to you.
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    You are solely responsible for ensuring that your use of the generated code complies with all applicable laws, regulations, and third-party licenses. DoPlan.dev does not guarantee that generated code is error-free, secure, or suitable for any specific purpose.
                                </p>
                            </section>

                            <section id="third-party" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">8. Third-Party Links</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    Our Service may contain links to third-party web sites or services that are not owned or controlled by DoPlan.dev. DoPlan.dev has no control over, and assumes no responsibility for, the content, privacy policies, or practices of any third party web sites or services.
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    You further acknowledge and agree that DoPlan.dev shall not be responsible or liable, directly or indirectly, for any damage or loss caused or alleged to be caused by or in connection with use of or reliance on any such content, goods or services available on or through any such web sites or services.
                                </p>
                            </section>

                            <section id="termination" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">9. Termination</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    We may terminate or suspend your account, without prior notice or liability, for any reason whatsoever, including without limitation if you breach the Terms. Upon termination, your right to use the Service will immediately cease.
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    If you wish to terminate your account, you may simply discontinue using the Service. All provisions of the Terms which by their nature should survive termination shall survive termination, including, without limitation, ownership provisions, warranty disclaimers, indemnity and limitations of liability.
                                </p>
                            </section>

                            <section id="indemnification" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">10. Indemnification</h2>
                                <p className="text-muted-foreground leading-relaxed">
                                    You agree to defend, indemnify and hold harmless DoPlan.dev and its licensee and licensors, and their employees, contractors, agents, officers and directors, from and against any and all claims, damages, obligations, losses, liabilities, costs or debt, and expenses (including but not limited to attorney's fees), resulting from or arising out of a) your use and access of the Service, by you or any person using your account and password, or b) a breach of these Terms.
                                </p>
                            </section>

                            <section id="limitation" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">11. Limitation of Liability</h2>
                                <p className="text-muted-foreground leading-relaxed">
                                    In no event shall DoPlan.dev, nor its directors, employees, partners, agents, suppliers, or affiliates, be liable for any indirect, incidental, special, consequential or punitive damages, including without limitation, loss of profits, data, use, goodwill, or other intangible losses, resulting from (i) your access to or use of or inability to access or use the Service; (ii) any conduct or content of any third party on the Service; (iii) any content obtained from the Service; and (iv) unauthorized access, use or alteration of your transmissions or content, whether based on warranty, contract, tort (including negligence) or any other legal theory, whether or not we have been informed of the possibility of such damage, and even if a remedy set forth herein is found to have failed of its essential purpose.
                                </p>
                            </section>

                            <section id="disclaimer" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">12. Disclaimer</h2>
                                <p className="text-muted-foreground leading-relaxed">
                                    Your use of the Service is at your sole risk. The Service is provided on an "AS IS" and "AS AVAILABLE" basis. The Service is provided without warranties of any kind, whether express or implied, including, but not limited to, implied warranties of merchantability, fitness for a particular purpose, non-infringement or course of performance.
                                </p>
                            </section>

                            <section id="governing-law" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">13. Governing Law</h2>
                                <p className="text-muted-foreground leading-relaxed">
                                    These Terms shall be governed and construed in accordance with the laws of the United States, without regard to its conflict of law provisions. Our failure to enforce any right or provision of these Terms will not be considered a waiver of those rights. If any provision of these Terms is held to be invalid or unenforceable by a court, the remaining provisions of these Terms will remain in effect.
                                </p>
                            </section>

                            <section id="changes" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">14. Changes to Terms</h2>
                                <p className="text-muted-foreground leading-relaxed">
                                    We reserve the right, at our sole discretion, to modify or replace these Terms at any time. If a revision is material we will try to provide at least 30 days notice prior to any new terms taking effect. What constitutes a material change will be determined at our sole discretion. By continuing to access or use our Service after those revisions become effective, you agree to be bound by the revised terms.
                                </p>
                            </section>

                            <section id="contact" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">15. Contact Us</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    If you have any questions about these Terms, please contact us:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li>By email: <a href="mailto:legal@doplan.dev" className="text-primary hover:underline">legal@doplan.dev</a></li>
                                    <li>By visiting this page on our website: <a href="https://doplan.dev/contact" className="text-primary hover:underline">https://doplan.dev/contact</a></li>
                                    <li>By mail: DoPlan Inc., 123 AI Boulevard, Tech City, TC 90210</li>
                                </ul>
                            </section>
                        </div>
                    </div>
                </div>
            </main>
            <Footer />
        </div>
    );
}
