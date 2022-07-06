##### Summary

<!--- Explain the objectives of this pull request. Don't forget to add Linear task number. -->
LINEAR:

##### Which modules will your code affect?

<!--- Which features will affected by this pull request? -->

##### How to test

<!--- Explain the testing procedure to the tester. Fill in as much as you can. -->

##### Test Spec

<!--- Fill in complete test cases. You might need to consult QA for additional test cases.
      Each case should always has a backup unit test.
      Unit tests should be readable along with the cases.-->

- [ ] ex. Case A...

##### Checklists

<!--- Use following checklist to check your code and check them before having other review -->

- [ ] Unit test every code paths.
- [ ] Locally test for happy and edge cases.
- [ ] No duplicate code... if so refactor now.
- [ ] No deprecated class/method.
- [ ] No valid username/password/secret/key in any parts of this project
- [ ] Database query
    - [ ] N/A
    - [ ] Already has index defined
    - [ ] No slowlog
- [ ] API
    - [ ] N/A
    - [ ] Added API change in api-changes document/page
    - [ ] Support Batch API for reducing number of requests to minimum
    - [ ] Support Pagination
- [ ] remote API call
    - [ ] N/A
    - [ ] Calling using Batch API for reducing number of requests to minimum
- [ ] default value / overloading
    - [ ] N/A
    - [ ] Backward compatibility
- [ ] config
    - [ ] N/A
    - [ ] Not expose the secret in config file
- [ ] log kv
    - [ ] N/A
    - [ ] Key is snakecase
    - [ ] Ensure that there is no different data type of value of the same key to prevent log parser failure
- [ ] Cherry-pick the commit to dev , prepared data and tested