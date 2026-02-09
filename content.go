package main

// Profile holds the top-level bio information.
type Profile struct {
	Name     string
	Role     string
	Location string
	Bio      string
}

// Experience represents a single work history entry.
type Experience struct {
	Title       string
	Company     string
	Period      string
	Description string
	Highlights  []string
}

// Project represents a portfolio project.
type Project struct {
	Name        string
	Description string
	Tech        []string
	URL         string
}

// SkillGroup is a named category of skills.
type SkillGroup struct {
	Category string
	Skills   []string
}

// Education represents a degree or certification.
type Education struct {
	Degree      string
	Institution string
	Period      string
	Details     string
}

// ContactInfo holds a single contact method.
type ContactInfo struct {
	Label string
	Value string
}

// --- Portfolio data ---

var profile = Profile{
	Name:     "Daniel Vaughan",
	Role:     "Senior Software Engineer",
	Location: "West Lafayette, IN",
	Bio: `Senior software engineer specializing in backend infrastructure,
cloud-native systems, and developer tooling. I build RESTful APIs,
design scalable infrastructure on AWS, and containerize everything
with Kubernetes and Docker. Skilled at translating complex technical
workflows for non-technical audiences and collaborating across teams
to ship reliable, maintainable software.`,
}

var experiences = []Experience{
	{
		Title:       "Senior Software Engineer",
		Company:     "Inari Agriculture",
		Period:      "Jun 2021 — Present",
		Description: "Backend infrastructure, API development, and cloud platform engineering.",
		Highlights: []string{
			"Develops RESTful APIs utilizing Flask, SQLAlchemy, PostgreSQL, Celery, and Redis for streamlined computational jobs and data access",
			"Containerizes API servers and Celery workers for deployment on Kubernetes (EKS) with Helm",
			"Utilizes AWS services (RDS, EKS, ECR, S3) to increase service scalability and availability",
			"Develops CI/CD workflows with CircleCI for quick turnaround during development",
			"Designs CLI tools and API endpoints using Argo for job orchestration on Kubernetes",
			"Architects per-account infrastructure using Terraform for state management and separation of sensitive data",
			"Worked with external consultants on IaC modules for cross-region and cross-account data redundancy",
		},
	},
	{
		Title:       "Senior Analyst & Developer",
		Company:     "RSA AFCC (Dell-EMC subsidiary)",
		Period:      "Jan 2019 — May 2021",
		Description: "Internal tooling and international anti-fraud operations.",
		Highlights: []string{
			"Developed web-based internal tools for quality assurance with full SQL database connectivity and email reports",
			"Worked as part of an international team, contacting CERTs and hosting providers worldwide to meet client needs",
		},
	},
	{
		Title:       "Senior Capstone Project — Minecraft",
		Company:     "Mojang Studios",
		Period:      "Aug 2020 — Dec 2020",
		Description: "Collaborated with Mojang on Minecraft's save system.",
		Highlights: []string{
			"Researched plausible updates to Minecraft's save system for more efficient data access and resilience against data loss",
			"Began implementations of save system updates in C++ for utilization in further development",
		},
	},
}

var projects = []Project{
	{
		Name:        "ssh-portfolio",
		Description: "This very site! A terminal-based portfolio served over SSH, built with Go and the Charm ecosystem.",
		Tech:        []string{"Go", "Wish", "Bubble Tea", "Lip Gloss"},
		URL:         "github.com/vaughancodes/ssh-portfolio",
	},
	{
		Name:        "web-portfolio",
		Description: "A React-based terminal-themed portfolio inspired by the SSH version.",
		Tech:        []string{"React", "TypeScript", "Vite"},
		URL:         "github.com/vaughancodes/web-portfolio",
	},
}

var skillGroups = []SkillGroup{
	{
		Category: "Languages",
		Skills:   []string{"Python", "C/C++", "SQL", "Bash"},
	},
	{
		Category: "Cloud & Infrastructure",
		Skills:   []string{"AWS (EC2, EKS, RDS, ECR, S3)", "Kubernetes", "Docker", "Terraform", "Helm"},
	},
	{
		Category: "Tools & Platforms",
		Skills:   []string{"Argo", "CircleCI", "Sysdig", "ELK", "Git"},
	},
	{
		Category: "Frameworks & Libraries",
		Skills:   []string{"Flask", "SQLAlchemy", "Celery", "Redis", "PostgreSQL"},
	},
	{
		Category: "Soft Skills",
		Skills:   []string{"Technical Communication", "Stakeholder Collaboration", "Cross-team Coordination"},
	},
}

var education = []Education{
	{
		Degree:      "B.S. Computer Science",
		Institution: "Purdue University, West Lafayette, IN",
		Period:      "May 2021",
		Details:     "Focus in Software Engineering and Security.",
	},
}

var contacts = []ContactInfo{
	{Label: "Email", Value: "daniel@vaughan.codes"},
	{Label: "Phone", Value: "(336) 380-3600"},
	{Label: "Office", Value: "(765) 201-4560"},
	{Label: "Portfolio", Value: "vaughan.codes"},
	{Label: "GitHub", Value: "github.com/vaughancodes"},
	{Label: "LinkedIn", Value: "linkedin.com/in/vaughancodes"},
	{Label: "Location", Value: "West Lafayette, IN"},
}
