async function handler (event, context) {
    // TODO: Handle message...
    const records = event.Records
    console.log("extra log")
    console.log(records)
    
    return {}
  }
  
  module.exports.handler = handler