---
name: optimize
category: optimize
trigger: "/optimize [<subcommand>]"
description: "Project optimization hub"
agentInvolvement:
  - Design & UX Manager
  - UI/UX Designer
  - DevOps Engineer
  - Performance Engineer
filesRead:
  - ".do/system/DESIGN_SYSTEM.md"
  - ".do/system/ARCHITECTURE.md"
  - "src/**"
filesModified:
  - ".do/system/DESIGN_SYSTEM.md"
  - ".do/system/COST_OPTIMIZATION.md"
  - ".do/system/PERFORMANCE_OPTIMIZATION.md"
examples:
  - "/optimize → Show optimization menu"
  - "/optimize design → UI/UX improvements"
  - "/optimize finance → Cost optimization"
  - "/optimize performance → Performance optimization"
  - "/optimize all → Run all optimizations"
---

When user types /optimize or /optimize <subcommand>:

1. **If no subcommand provided**: Show menu:
   - "Select optimization type:"
   - "1. Design - UI/UX improvements"
   - "2. Finance - Cost optimization"
   - "3. Performance - Performance optimization"
   - "4. All - Run all optimizations"
   - Wait for user selection

2. **Subcommand: design** (or user selects option 1):
   - UI/UX Review: Design Manager and UI/UX Designer review interface
   - Improvement Suggestions: Provide UI/UX improvement recommendations
   - Update Design System: Update DESIGN_SYSTEM.md with improvements
   - Response: "UI/UX improvements suggested! Review design updates in DESIGN_SYSTEM.md"

3. **Subcommand: finance** (or user selects option 2):
   - Cost Analysis: Analyze current infrastructure and service costs
   - Optimization Recommendations: Provide cost optimization suggestions
   - Generate Cost Plan: Create cost optimization plan
   - Response: "Cost analysis complete! Review optimization recommendations in COST_OPTIMIZATION.md"

4. **Subcommand: performance** (or user selects option 3):
   - Performance Analysis: Performance Engineer analyzes application performance
   - Identify Bottlenecks: Detect performance bottlenecks and issues
   - Optimization Suggestions: Provide performance optimization recommendations
   - Generate Performance Report: Create performance optimization report
   - Response: "Performance analysis complete! Review optimization recommendations in PERFORMANCE_OPTIMIZATION.md"

5. **Subcommand: all** (or user selects option 4):
   - Run all three optimizations in sequence (design, finance, performance)
   - Generate comprehensive optimization report
   - Response: "All optimizations complete! Review reports in DESIGN_SYSTEM.md, COST_OPTIMIZATION.md, and PERFORMANCE_OPTIMIZATION.md"

