import Header from './Header';
import Hero from './Hero';
import Features from './Features';
import Workflow from './Workflow';
import Footer from './Footer';
import SEO from './SEO';

export default function Home() {
  return (
    <div className="min-h-screen bg-background">
      <SEO
        title="Home"
        description="DoPlan - The Zero-Install AI Project Director that orchestrates multiple AI agents to handle your entire development workflow."
        canonical="/"
      />
      <Header />
      <main>
        <Hero />
        <Workflow />
        <Features />
      </main>
      <Footer />
    </div>
  );
}
