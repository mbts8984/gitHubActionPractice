module.exports = ({github, context}) => {
    console.log("MADE IT TO SCRIPTS print")
    console.log('Context print ', context)
    return context.payload.client_payload.value
  }