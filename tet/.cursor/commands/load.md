# /load

## Trigger
/load <path>

## Examples
- /load @library/04-frameworks/frontend/nextjs.md
- /load .plan/00_System/PRD.md


## Action
When user types /load <path>:

1. **Parse Path**: Extract file or directory path
2. **Load Content**: Read the specified file or files from directory
3. **Inject Context**: Add content to agent context for current session
4. **Response**: "Context loaded! [path] is now available to agents."

## Agent Involvement
- **Project Orchestrator**

## Files Read
- .cursor/rules/library/**
- .plan/**

