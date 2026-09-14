1. Scenario 1: Platform vs Product team
- Data layer migration
	- Discussions about using right platfom to use for the data layer.
	- Push from product team in one direction (create their own db system for faster delivery) vs infrastructure (risk of fragmentation, security,performance, long term maintenance)

2. Technical Debt vs Feature Velocity
	1. Move to new IdP, replace data layer, make foundational gateway
		1. Postpone the foundational gateway work and worked with non foundational gateway for the initial phases in order to allocate more resources to the data layer replacement and cut less corners.
		2. Take advantage of SRV DNS records to avoid gateway.
			1. Things added back: Required updates to the webstore infra to work in our scenarios. Changes to use a regional model vs a global idp.
			2. Proper test framework investments vs relying on third party test beds.
			3. Reduced risk by having additional resources from the beginning. 
3. Interpersonal Clashes Between Senior Engineers
	1. UX framework used for login page
		1. Custom in-house framework, very light weight. Used in multiple projects. Hard to develop, preview changes.
		2. Vs. New Open source based framework.  Webpack for dependency tracking, linting pipeline, mvvm framework for UX binding (knockout), UX in html made transition from prototype/spec to final product. 
		3. Result: marginally bigger files, mitigated with cached common files. Greatly increased developer productivity.
4. Scope Creep and Stakeholder Alignment
	1. Attempt to include WS-Fed/WS-Trust support as part of the migration to the new IdP (dSTS)
		1. Original plan included enabling the new IdP for the OAuth protocol