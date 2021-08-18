module.exports = async ({github, context}) => {
    console.log("MADE IT TO SCRIPTS print")
    // console.log('Context print ', context)

    const matchesPattern = (pattern, text) => {
        const regex = new RegExp(pattern);
        return !!text.match(regex)
      };
    
    const files = await context.github.pulls.listFiles({
        owner,
        repo,
        number: pullNumber,
        mediaType: {
          format: "diff",
        },
    });
    console.log("DIFF1: ", files)

    const diff_url = context.payload.pull_request.diff_url
    const result = await github.request(diff_url)
    console.log("DIFF2: ", result)


    github.issues.createComment({
        issue_number: context.issue.number,
        owner: context.repo.owner,
        repo: context.repo.repo,
        body: '👋 Thanks for reporting!'
      })
    
    
  }