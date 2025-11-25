"use client";

import React, { useState } from "react";
import Link from "next/link";

export function Header() {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  return (
    <header className="w-full border-b border-border bg-bg-primary">
      <nav className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          {/* Logo */}
          <Link
            href="/"
            className="text-xl font-bold text-primary hover:text-primary-dark transition-colors"
          >
            QR Generator
          </Link>

          {/* Desktop Navigation */}
          <div className="hidden md:flex items-center gap-6">
            <Link
              href="/#api-docs"
              className="text-sm font-medium text-text-secondary hover:text-text-primary transition-colors"
            >
              API Docs
            </Link>
            <Link
              href="/#analytics"
              className="text-sm font-medium text-text-secondary hover:text-text-primary transition-colors"
            >
              Analytics
            </Link>
          </div>

          {/* Mobile Menu Button */}
          <button
            type="button"
            className="md:hidden p-2 text-text-secondary hover:text-text-primary"
            onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
            aria-label="Toggle menu"
            aria-expanded={mobileMenuOpen}
          >
            <svg
              className="w-6 h-6"
              fill="none"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              {mobileMenuOpen ? (
                <path d="M6 18L18 6M6 6l12 12" />
              ) : (
                <path d="M4 6h16M4 12h16M4 18h16" />
              )}
            </svg>
          </button>
        </div>

        {/* Mobile Menu */}
        {mobileMenuOpen && (
          <div className="md:hidden py-4 border-t border-border">
            <Link
              href="/#api-docs"
              className="block py-2 text-sm font-medium text-text-secondary hover:text-text-primary transition-colors"
              onClick={() => setMobileMenuOpen(false)}
            >
              API Docs
            </Link>
            <Link
              href="/#analytics"
              className="block py-2 text-sm font-medium text-text-secondary hover:text-text-primary transition-colors"
              onClick={() => setMobileMenuOpen(false)}
            >
              Analytics
            </Link>
          </div>
        )}
      </nav>
    </header>
  );
}

