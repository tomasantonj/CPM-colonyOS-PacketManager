# CPM Dashboard Development Tickets
Based on: [CPM_Dashboard_App_Requirements.md](./CPM_Dashboard_App_Requirements.md)

## 1. Project Setup & Core Logic

### [FRONTEND-01] Initialize React/Vite Project
**Priority:** High
**Description:**
Scaffold the application using the defined tech stack.
**Technical Details:**
- Vite + React + TypeScript.
- Setup Tailwind CSS & Shadcn/UI.
- Configure Directories (`src/features`, `src/components`, etc.).
**Acceptance Criteria:**
- App runs locally (`npm run dev`).
- Basic "Hello World" with Tailwind styling works.

### [FRONTEND-02] Global Layout & Navigation
**Priority:** High
**Description:**
Create the main application shell.
**Technical Details:**
- Sidebar Navigation (Collapsible).
- Dark Mode toggle/logic (Default: Dark).
- React Router setup (`/`, `/registry`, `/builder`, `/settings`).
**Acceptance Criteria:**
- User can navigate between pages.
- Layout is responsive (Mobile/Desktop).

### [FRONTEND-03] Authentication Logic
**Priority:** High
**Description:**
Implement Login flows for Registry and ColonyOS.
**Technical Details:**
- **Registry Auth:** Login form -> API -> JWT storage.
- **ColonyOS Auth:** Input for Private Key (SessionStorage only).
**Acceptance Criteria:**
- User can log in.
- "Protected Routes" redirect to login if no token.

## 2. Feature: Registry Browser

### [FRONTEND-04] Registry Package Grid
**Priority:** High
**Description:**
View searching and listing of available packages.
**Technical Details:**
- Fetch data from API (or Mock).
- Render "Package Cards" (Name, Avatar, Description).
- Implement Search/Filter bar.
**Acceptance Criteria:**
- Grid displays mocked/real data.
- Search filters the list client-side or server-side.

### [FRONTEND-05] Package Detail View
**Priority:** Medium
**Description:**
View full details of a specific package.
**Technical Details:**
- Route: `/registry/:packageName`.
- Render `README.md` using a Markdown renderer.
- Display `inputs` list.
**Acceptance Criteria:**
- Users can read package documentation within the app.

## 3. Feature: Visual Builder

### [FRONTEND-06] Workflow Canvas (React Flow)
**Priority:** Critical
**Description:**
The drag-and-drop workspace for building colonies.
**Technical Details:**
- Integrate React Flow.
- Implement "Drag from Palette" to "Drop on Canvas".
**Acceptance Criteria:**
- User can drop nodes onto the canvas.
- Nodes are connected visually.

### [FRONTEND-07] Node Configuration Panel
**Priority:** Critical
**Description:**
Dynamic form generation based on Package Inputs.
**Technical Details:**
- When a node is selected, show a side panel.
- Loop through `spec.inputs`.
- Render corresponding Input/Select/Switch components.
- Validate required fields.
**Acceptance Criteria:**
- Selecting a node shows the correct form.
- Form updates the node's internal state.

### [FRONTEND-08] YAML Preview & Export
**Priority:** High
**Description:**
Real-time generation of the `colony.yaml` file.
**Technical Details:**
- Convert React Flow state + Node Form Data -> JSON/YAML.
- Display in a syntax-highlighted code block (Right or Bottom pane).
- "Download YAML" button.
**Acceptance Criteria:**
- Changes in the canvas update the YAML text immediately.
- Generated YAML is valid ColonyOS syntax.

## 4. Polish & Docs

### [FRONTEND-09] In-App Documentation
**Priority:** Low
**Description:**
Help users understand how to use the tool.
**Technical Details:**
- "Getting Started" overlay.
- Tooltips on complex terms.
**Acceptance Criteria:**
- First-time visit shows the guide.
