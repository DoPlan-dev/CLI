"use client";

import React, { useState, useEffect } from "react";
import { Analytics } from "@/types/analytics";

export function AnalyticsDisplay({ className = "" }: { className?: string }) {
  const [analytics, setAnalytics] = useState<Analytics | null>(null);
  const [loading, setLoading] = useState(true);

  const fetchAnalytics = async () => {
    try {
      const response = await fetch("/api/analytics");
      if (response.ok) {
        const data: Analytics = await response.json();
        setAnalytics(data);
      }
    } catch (error) {
      console.error("Failed to fetch analytics:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAnalytics();
    // Auto-refresh every 30 seconds
    const interval = setInterval(fetchAnalytics, 30000);
    return () => clearInterval(interval);
  }, []);

  if (loading || !analytics) {
    return (
      <div
        className={`flex items-center gap-4 p-4 bg-bg-secondary rounded-lg ${className}`}
      >
        <div className="text-text-secondary text-sm">Loading analytics...</div>
      </div>
    );
  }

  const getStatusColor = () => {
    switch (analytics.apiStatus) {
      case "operational":
        return "bg-success";
      case "degraded":
        return "bg-warning";
      case "down":
        return "bg-error";
      default:
        return "bg-text-tertiary";
    }
  };

  return (
    <div
      className={`flex items-center gap-4 p-4 bg-bg-secondary rounded-lg ${className}`}
    >
      <div className="flex flex-col gap-1">
        <span className="text-xs text-text-secondary uppercase tracking-wide">
          Total Generations
        </span>
        <span className="text-2xl font-bold text-primary">
          {analytics.totalGenerations.toLocaleString()}
        </span>
      </div>
      <div className="flex flex-col gap-1">
        <span className="text-xs text-text-secondary uppercase tracking-wide">
          Today
        </span>
        <span className="text-2xl font-bold text-primary">
          {analytics.todayGenerations.toLocaleString()}
        </span>
      </div>
      <div className="flex items-center gap-2 ml-auto">
        <div
          className={`w-3 h-3 rounded-full ${getStatusColor()} pulse`}
          aria-label={`API Status: ${analytics.apiStatus}`}
        />
        <span className="text-sm text-text-secondary capitalize">
          {analytics.apiStatus}
        </span>
      </div>
    </div>
  );
}

