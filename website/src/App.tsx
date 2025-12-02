import { Suspense, lazy } from "react";
import { Routes, Route } from "react-router-dom";
import { HelmetProvider } from 'react-helmet-async';
import { ThemeProvider } from "@/hooks/useTheme";
import CookieConsent from "./components/CookieConsent";

const Home = lazy(() => import("./components/home"));
const FeaturesPage = lazy(() => import("./components/FeaturesPage"));
const PrivacyPolicy = lazy(() => import("./components/PrivacyPolicy"));
const TermsOfService = lazy(() => import("./components/TermsOfService"));
const CommandsPage = lazy(() => import("./components/CommandsPage"));
const DocsPage = lazy(() => import("./components/docs/DocsPage"));

function App() {
  return (
    <HelmetProvider>
      <ThemeProvider>
        <Suspense fallback={<p>Loading...</p>}>
          <>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/commands" element={<CommandsPage />} />
              <Route path="/features" element={<FeaturesPage />} />
              <Route path="/docs" element={<DocsPage />} />
              <Route path="/docs/:section/:page" element={<DocsPage />} />
              <Route path="/privacy" element={<PrivacyPolicy />} />
              <Route path="/terms" element={<TermsOfService />} />
            </Routes>
            <CookieConsent />
          </>
        </Suspense>
      </ThemeProvider>
    </HelmetProvider>
  );
}

export default App;
