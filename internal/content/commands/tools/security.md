---
name: security
category: tools
trigger: "/security [<subcommand>]"
description: "Security review and audit"
agentInvolvement:
  - Security Lead
filesRead:
  - "src/**"
  - ".do/system/ARCHITECTURE.md"
  - ".do/system/SECURITY.md"
filesModified:
  - ".do/system/SECURITY.md"
  - ".do/system/SECURITY_AUDIT.md"
examples:
  - "/security → Show security menu"
  - "/security review → Security review"
  - "/security audit → Security audit"
  - "/security both → Review and audit"
---

When user types /security or /security <subcommand>:

1. **If no subcommand provided**: Show menu:
   - "Select security operation:"
   - "1. Review - Security review"
   - "2. Audit - Security audit"
   - "3. Both - Review and audit"
   - Wait for user selection

2. **Subcommand: review** (or user selects option 1):
   - Security Review: Security Lead conducts security review
   - Vulnerability Assessment: Identify and document security vulnerabilities
   - Generate Report: Create security review report
   - Response: "Security review complete! Review security findings in SECURITY.md"

3. **Subcommand: audit** (or user selects option 2):
   - Security Audit: Security Lead conducts comprehensive security audit
   - Vulnerability Scanning: Scan for security vulnerabilities
   - Compliance Check: Verify compliance with security standards
   - Generate Audit Report: Create security audit report
   - Response: "Security audit complete! Review audit findings in SECURITY_AUDIT.md"

4. **Subcommand: both** (or user selects option 3):
   - Run both review and audit in sequence
   - Generate comprehensive security report
   - Response: "Security review and audit complete! Review findings in SECURITY.md and SECURITY_AUDIT.md"

