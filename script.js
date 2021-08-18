module.exports = async ({github, context}) => {
    console.log("MADE IT TO SCRIPTS print")
    // console.log('Context print ', context)

    const matchesPattern = (pattern, text) => {
        const regex = new RegExp(pattern);
        return !!text.match(regex)
      };

    const diff_url = context.payload.pull_request.diff_url
    const result = await github.request(diff_url)
    console.log("DIFF2: ", result)


    github.issues.createComment({
        issue_number: context.issue.number,
        owner: context.repo.owner,
        repo: context.repo.repo,
        body: '👋 Thanks for reporting!'
      })
    
    try {
        const attemptOne = await github.pull_request.createComment({
            body: 'dogPop, dogPop, dogPop',
            owner: context.repo.owner,
            repo: context.repo.repo,
            line: 7,
            path: "README.md",
            pull_number: context.pull_request.number,
            commit_id: context.pull_request.commit_id
        })
    } catch (error) {
        console.log("error", error)
    }
  }