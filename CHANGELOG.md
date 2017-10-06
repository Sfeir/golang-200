# Changelog
## v0.0.2
- 17/10/05 fix(test): enhance web test errors on full server test (SFR)
           refact(model): refact model constructor
- 17/10/04 chore(make): clean make file auto generated help (SFR for RLE)
           chore(dep): update package manager to dep, update vendor
           chore(docker): update docker image version, add multistep build
           refact(main): refact arg parsing, remove redondant arg parsing
- 17/04/14 doc(stat): clarify PlusOne go doc (SFR)
           doc(model): fix the NewID documentation to match method content
## v0.0.1 [17/03/21]
- 17/03/21 fix(main): fix the statd flag help message (SFR)
- 17/03/20 fix(web): fix the Fatal log level in server build (SFR)
- 17/03/17 refact(web): rename handler to controller (SFR)
           chore(test): add postman collection for testing
           chore(doc): update vendor manager in README
           refact(dao): interface type safety check, json omitempty tag
           chore(make): make the bench tool work again
           refact(web): update web test to use MongoDB
           fix(main): fix timezone for bson/json marshalling
- 17/03/14 refact(dao): go fmt src, update DAO with typed const (SFR)
           chore(make): clean the clean target to prevent glide.lock removal
- 17/03/12 refact(dao): update model and dao to use UUID (SFR)
           chore(doc): update main architecture diagram
- 17/02/21 chore(vendor): update vendors to glide (SFR)
           refact(model): add const types
           chore(etc): update scripts and types to todos
- 16/12/20 chore(all): refactor from handsongo to todolost (OFU)

