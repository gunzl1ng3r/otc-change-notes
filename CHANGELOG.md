# CHANGELOG

## 2019-09-17
* REQ002 has been fulfilled
  * `getPages` retrieves all the pages linked at a given page
  * `identifyChangeHistory` takes a links and looks for the title `Change History`
  * `main` call `getPages`, sends the retrieved links to `identifyChangeHistory` and stores identified pages in a new object
  * links of pages containing a Change History are then evaluated using `parseChangeHistory`

## 2019-05-20
* REQ001 has been fulfilled
  * given a static link to a "Change History" page, this tool will turn that page into a map and print it
