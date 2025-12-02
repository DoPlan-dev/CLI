import { useEffect, useState } from 'react';
import Header from './Header';
import Footer from './Footer';
import { cn } from '@/lib/utils';
import SEO from './SEO';

const sections = [
    { id: 'introduction', title: '1. Introduction' },
    { id: 'information-collection', title: '2. Information Collection' },
    { id: 'information-usage', title: '3. How We Use Information' },
    { id: 'cookies', title: '4. Cookies & Tracking' },
    { id: 'data-retention', title: '5. Data Retention' },
    { id: 'data-transfer', title: '6. Data Transfer' },
    { id: 'data-disclosure', title: '7. Disclosure of Data' },
    { id: 'data-security', title: '8. Data Security' },
    { id: 'user-rights', title: '9. User Rights' },
    { id: 'children', title: '10. Children\'s Privacy' },
    { id: 'changes', title: '11. Changes to Policy' },
    { id: 'contact', title: '12. Contact Us' },
];

export default function PrivacyPolicy() {
    const [activeSection, setActiveSection] = useState('introduction');

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
                title="Privacy Policy"
                description="DoPlan Privacy Policy - How we collect, use, and protect your data."
                canonical="/privacy"
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
                        <h1 className="text-4xl sm:text-5xl font-bold mb-6">Privacy Policy</h1>
                        <p className="text-muted-foreground mb-12 text-lg">
                            Last updated: {new Date().toLocaleDateString()}
                        </p>

                        <div className="prose prose-lg dark:prose-invert max-w-none space-y-16">
                            <section id="introduction" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">1. Introduction</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    Welcome to DoPlan.dev ("we," "our," or "us"). We are committed to protecting your personal information and your right to privacy. This Privacy Policy explains in detail how we collect, use, disclose, and safeguard your information when you use our website, CLI tools, API, and related services (collectively, the "Service").
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    We take your privacy seriously. By accessing or using our Service, you acknowledge that you have read, understood, and agree to be bound by the terms of this Privacy Policy. If you do not agree with our policies and practices, please do not use our Service.
                                </p>
                            </section>

                            <section id="information-collection" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">2. Information Collection</h2>
                                <p className="text-muted-foreground leading-relaxed mb-6">
                                    We collect several different types of information for various purposes to provide and improve our Service to you.
                                </p>

                                <h3 className="text-2xl font-semibold mb-4">2.1 Personal Data</h3>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    While using our Service, we may ask you to provide us with certain personally identifiable information that can be used to contact or identify you ("Personal Data"). Personally identifiable information may include, but is not limited to:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3 mb-6">
                                    <li><strong>Identity Information:</strong> First name, last name, username, or similar identifier.</li>
                                    <li><strong>Contact Information:</strong> Email address, telephone number, billing address, and shipping address.</li>
                                    <li><strong>Professional Information:</strong> Job title, company name, and industry.</li>
                                    <li><strong>Payment Information:</strong> Credit card details and billing information (processed securely by our third-party payment processors).</li>
                                </ul>

                                <h3 className="text-2xl font-semibold mb-4">2.2 Usage Data</h3>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    We may also collect information that your browser sends whenever you visit our Service or when you access the Service by or through a mobile device or CLI tool ("Usage Data"). This Usage Data may include information such as:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li>Internet Protocol (IP) address</li>
                                    <li>Browser type and browser version</li>
                                    <li>Operating system and platform</li>
                                    <li>The pages of our Service that you visit</li>
                                    <li>The time and date of your visit</li>
                                    <li>The time spent on those pages</li>
                                    <li>Unique device identifiers and other diagnostic data</li>
                                    <li>CLI command execution statistics (anonymized)</li>
                                    <li>Error logs and crash reports</li>
                                </ul>
                            </section>

                            <section id="information-usage" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">3. How We Use Information</h2>
                                <p className="text-muted-foreground leading-relaxed mb-6">
                                    DoPlan.dev uses the collected data for various specific purposes:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li><strong>Service Provision:</strong> To provide, operate, and maintain our Service, including the CLI tool and web dashboard.</li>
                                    <li><strong>Communication:</strong> To notify you about changes to our Service, security updates, and administrative messages.</li>
                                    <li><strong>Customer Support:</strong> To provide customer support and respond to your inquiries.</li>
                                    <li><strong>Improvement:</strong> To gather analysis or valuable information so that we can improve our Service, develop new features, and enhance user experience.</li>
                                    <li><strong>Monitoring:</strong> To monitor the usage of our Service and detect, prevent, and address technical issues.</li>
                                    <li><strong>Marketing:</strong> To provide you with news, special offers, and general information about other goods, services, and events which we offer that are similar to those that you have already purchased or enquired about unless you have opted not to receive such information.</li>
                                    <li><strong>Legal Compliance:</strong> To comply with applicable laws, regulations, and legal processes.</li>
                                </ul>
                            </section>

                            <section id="cookies" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">4. Cookies & Tracking Technologies</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    We use cookies and similar tracking technologies to track the activity on our Service and hold certain information. Cookies are files with small amount of data which may include an anonymous unique identifier.
                                </p>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    Tracking technologies also used are beacons, tags, and scripts to collect and track information and to improve and analyze our Service. You can instruct your browser to refuse all cookies or to indicate when a cookie is being sent. However, if you do not accept cookies, you may not be able to use some portions of our Service.
                                </p>
                                <h3 className="text-2xl font-semibold mb-4">Types of Cookies We Use:</h3>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li><strong>Essential Cookies:</strong> Necessary for the operation of the website (e.g., login sessions).</li>
                                    <li><strong>Preference Cookies:</strong> Allow us to remember your preferences (e.g., language, theme).</li>
                                    <li><strong>Analytics Cookies:</strong> Help us understand how visitors interact with the website.</li>
                                    <li><strong>Marketing Cookies:</strong> Used to track visitors across websites for marketing purposes.</li>
                                </ul>
                            </section>

                            <section id="data-retention" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">5. Data Retention</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    We will retain your Personal Data only for as long as is necessary for the purposes set out in this Privacy Policy. We will retain and use your Personal Data to the extent necessary to comply with our legal obligations (for example, if we are required to retain your data to comply with applicable laws), resolve disputes, and enforce our legal agreements and policies.
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    We will also retain Usage Data for internal analysis purposes. Usage Data is generally retained for a shorter period of time, except when this data is used to strengthen the security or to improve the functionality of our Service, or we are legally obligated to retain this data for longer time periods.
                                </p>
                            </section>

                            <section id="data-transfer" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">6. Data Transfer</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    Your information, including Personal Data, may be transferred to — and maintained on — computers located outside of your state, province, country or other governmental jurisdiction where the data protection laws may differ than those from your jurisdiction.
                                </p>
                                <p className="text-muted-foreground leading-relaxed">
                                    DoPlan.dev will take all steps reasonably necessary to ensure that your data is treated securely and in accordance with this Privacy Policy and no transfer of your Personal Data will take place to an organization or a country unless there are adequate controls in place including the security of your data and other personal information.
                                </p>
                            </section>

                            <section id="data-disclosure" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">7. Disclosure of Data</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    We may disclose your Personal Data in the good faith belief that such action is necessary to:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li><strong>Legal Requirements:</strong> To comply with a legal obligation or court order.</li>
                                    <li><strong>Rights Protection:</strong> To protect and defend the rights or property of DoPlan.dev.</li>
                                    <li><strong>Safety:</strong> To prevent or investigate possible wrongdoing in connection with the Service or to protect the personal safety of users of the Service or the public.</li>
                                    <li><strong>Liability:</strong> To protect against legal liability.</li>
                                    <li><strong>Business Transaction:</strong> If DoPlan.dev is involved in a merger, acquisition, or asset sale, your Personal Data may be transferred.</li>
                                </ul>
                            </section>

                            <section id="data-security" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">8. Data Security</h2>
                                <p className="text-muted-foreground leading-relaxed">
                                    The security of your data is important to us, but remember that no method of transmission over the Internet, or method of electronic storage is 100% secure. We use commercially acceptable means to protect your Personal Data, including encryption, firewalls, and secure socket layer (SSL) technology. However, we cannot guarantee its absolute security.
                                </p>
                            </section>

                            <section id="user-rights" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">9. Your Data Protection Rights</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    Depending on your location (e.g., EEA, California), you may have certain rights regarding your personal information:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li><strong>Access:</strong> The right to access, update or to delete the information we have on you.</li>
                                    <li><strong>Rectification:</strong> The right to have your information rectified if that information is inaccurate or incomplete.</li>
                                    <li><strong>Objection:</strong> The right to object to our processing of your Personal Data.</li>
                                    <li><strong>Restriction:</strong> The right to request that we restrict the processing of your personal information.</li>
                                    <li><strong>Portability:</strong> The right to be provided with a copy of the information we have on you in a structured, machine-readable and commonly used format.</li>
                                    <li><strong>Withdraw Consent:</strong> The right to withdraw your consent at any time where DoPlan.dev relied on your consent to process your personal information.</li>
                                </ul>
                            </section>

                            <section id="children" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">10. Children's Privacy</h2>
                                <p className="text-muted-foreground leading-relaxed">
                                    Our Service does not address anyone under the age of 18 ("Children"). We do not knowingly collect personally identifiable information from anyone under the age of 18. If you are a parent or guardian and you are aware that your Children has provided us with Personal Data, please contact us. If we become aware that we have collected Personal Data from children without verification of parental consent, we take steps to remove that information from our servers.
                                </p>
                            </section>

                            <section id="changes" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">11. Changes to This Privacy Policy</h2>
                                <p className="text-muted-foreground leading-relaxed">
                                    We may update our Privacy Policy from time to time. We will notify you of any changes by posting the new Privacy Policy on this page. We will let you know via email and/or a prominent notice on our Service, prior to the change becoming effective and update the "effective date" at the top of this Privacy Policy. You are advised to review this Privacy Policy periodically for any changes.
                                </p>
                            </section>

                            <section id="contact" className="scroll-mt-32">
                                <h2 className="text-3xl font-bold mb-6">12. Contact Us</h2>
                                <p className="text-muted-foreground leading-relaxed mb-4">
                                    If you have any questions about this Privacy Policy, please contact us:
                                </p>
                                <ul className="list-disc pl-6 text-muted-foreground space-y-3">
                                    <li>By email: <a href="mailto:privacy@doplan.dev" className="text-primary hover:underline">privacy@doplan.dev</a></li>
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