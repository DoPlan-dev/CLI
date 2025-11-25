import React from "react";
import Link from "next/link";

export function Footer() {
  const currentYear = new Date().getFullYear();

  return (
    <footer className="w-full border-t border-border bg-bg-secondary mt-24">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex flex-col md:flex-row items-center justify-between gap-4">
          <div className="text-sm text-text-secondary">
            © {currentYear} QR Generator. All rights reserved.
          </div>
          <div className="flex items-center gap-6">
            <Link
              href="https://github.com/DoPlan-dev/test-no01"
              target="_blank"
              rel="noopener noreferrer"
              className="text-sm text-text-secondary hover:text-text-primary transition-colors"
            >
              GitHub
            </Link>
            <Link
              href="/#api-docs"
              className="text-sm text-text-secondary hover:text-text-primary transition-colors"
            >
              API Docs
            </Link>
          </div>
        </div>
      </div>
    </footer>
  );
}

