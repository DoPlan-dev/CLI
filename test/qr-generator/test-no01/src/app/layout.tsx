import type { Metadata } from "next";
import "./globals.css";
import { Header } from "@/components/Header";
import { Footer } from "@/components/Footer";

export const metadata: Metadata = {
  title: "QR Code Generator API - Fast & Simple",
  description:
    "Generate QR codes in milliseconds. Fast, simple, and developer-friendly QR code generation API with PNG and SVG support.",
  keywords: ["QR code", "QR generator", "API", "QR code API", "barcode"],
  authors: [{ name: "DoPlan" }],
  openGraph: {
    title: "QR Code Generator API",
    description: "Generate QR codes in milliseconds",
    type: "website",
  },
  twitter: {
    card: "summary_large_image",
    title: "QR Code Generator API",
    description: "Generate QR codes in milliseconds",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <Header />
        {children}
        <Footer />
      </body>
    </html>
  );
}
