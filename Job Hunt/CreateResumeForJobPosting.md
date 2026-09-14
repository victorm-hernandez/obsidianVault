---
name: CreateResumeForJobPosting
---

# Instructions to create a resume based on a job posting

When working a job posting in the JobHunt project (`JobHunt/<UserName>/`), treat the documented steps as a checklist with explicit stop points, not as context to read once and improvise from.

## Purpose of this document

- Help the user tailor his resume (`JobHunt/<UserName>/Resumes/ResumeBase.md`) and create role-specific resumes and cover letters.
- You will help write a resume, and once it is done you will review it again from the perspective of a hiring manager, would this resume pass any automatic filtering? it will catch your atention? if not, go back and iterate on it.

## Inputs

### Required Inputs

If any of the inputs is not provided initially, ask the user for them.

- Username: The name of the user applying to a position, this usually will not change across sessions. If you need to ask, make sure to keep it in memory for future use. For example `VictorHernandez`
- Job Posting: Text containing the job posting or the URL you should use to retrieve it.

### Computed inputs

Only ask for these inputs if there is a problem calculating or finding them:

- Job Identifier (JobId): This value should be formed as follows `CompanyName-TeamName-YYYY-MM-DD`, if you cannot determine this from the job posting, ask the user. This is an example for the Company "Apple" and the team "Identity" and the current date is June 15th 2026 `Apple-Identity-2026-06-15` .

- Base resume: You will need the current user's base resume. This base resume is common to all the versions of the user's resume. This should be located under `JobHunt/<Username>/Resumes/ResumeBase.md`. For example `JobHunt/VictorHernandez/Resumes/ResumeBase.md`. If you cannot find it, ask the user for the right path.

- Resume addendums: Addendums are files that cover the information gaps of the base resume for type of role (canonical) or a specific position (job posting). You may use this as a knowledge base to learn from previous applications to assist on the current process. You may find the cannonical addendums under the path `JobHunt/<Usernane>/Resumes/Resume-<SpecialtyName>.md`. For example: `JobHunt/VictorHernandez/Resumes/Resume-FrontEndDev.md`. You may also find job posting specific addendums on the folders of previous applications: `JobHunt/<Username>/Applications/<JobId>/Resume-Addendum.md`. For example: `JobHunt/VictorHernandez/Applications/Apple-Identity-2026-06-15/Resume-Addendum.md`

- Previous project descriptions: You can find detailed descriptions of projects from the user on the following path: `BehavioralInterview/Projects/<Username>`, for example `BehavioralInterview/Projects/VictorHernandez`.

## Workflow

1. The user will provide you with a job posting (paste text or a link). If the job posting is a link, navigate the web to retrieve the job posting information.

2. You will make sure the proper folder for this posting exist under this path `JobHunt/<Username>/Applications/<JobIdentifier>/`. If the folder doesn't exists, create it. For example: `JobHunt/VictorHernandez/Applications/Apple-Identity-2026-06-15`.

3. Each posting should have a `Notes` file with a summary of the position. If it doesnt exist, you will create a job `Notes` file on the job posting folder, `JobHunt/<Username>/Applications/<JobIdentifier>/Notes.md` using the template defined on the following path `JobHunt/ApplicationNotesTemplate.md`

4. You will populate the `Notes` file with any relevant data not included on it already based on the data from the posting, do not generate or populate data from any other source but the job posting.

5. You will compare the user's base resume (`JobHunt/<Username>/Resumes/ResumeBase.md`) and note major gaps with the job posting.

6. You will read all the addendums available (canonical and job posting specific) and detailed project descriptions and use this information to propose addendums for missing items. The result of the analysis should be saved as `Resume-GapAnalisis.md` inside the `JobHunt/<Username>/Applications/<JobIdentifier>/` folder.

7. You will ask the user for each of the points you found and how the user suggest to close them.

8. With the information obtained you will create a new file with all of the new content that will be appended to the base resume (`JobHunt/<Username>/Resumes/ResumeBase.md`). This information will be used in future iterations to produce other resumes. You will save this file as `Resume-Addendum.md` inside the `JobHunt/<Username>/Applications/<JobIdentifier>/` folder.

9. Using the base (`JobHunt/<Username>/Resumes/ResumeBase.md`) and the addendum you just created (`Resume-Addendum.md`). You will produce a tailored resume `Resume.md` and cover letter `CoverLetter.md` in markdown and store it inside the `JobHunt/<Username>/Applications/<JobIdentifier>/` folder so everything for a specific posting is co-located.

10. You will ask the user if they would like to export it to PDF/HTML/Word.

## Output guidelines

- Do NOT include generation metadata in exported resumes (no timestamps, "Generated:" footers, or internal notes).
- Do NOT link to internal files or folders (e.g., "References & Addendum"). If an addendum contains useful examples or evidence, you will include those examples directly in the exported resume or cover letter so external viewers have necessary context.
- Keep exported documents professional and self-contained: all customer-facing artifacts must contain only the information a hiring manager needs to evaluate the application.
- When adding information from an addemdum, do not list them at the end of the document, try to integrate this as part of the job experience, to make it feel organic.
- Never use '—' since those are a clear indicator of AI.
- Do NOT include a "Languages" section in resumes.

These rules ensure exports are clean, self-contained, and safe to share with external reviewers.