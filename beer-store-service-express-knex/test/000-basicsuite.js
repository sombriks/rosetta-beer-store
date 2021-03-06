// cross-env NODE_ENV=testing knex migrate:latest ; cross-env NODE_ENV=testing mocha --timeout=30000 --exit
const chai = require('chai')

const { errfn, environment } = require("../src/components/config")

chai.should()

describe("Basic tests suite", _ => {

  // mocking the express response object
  const res = { status: _ => ({ send: _ => _ }) }

  it("should cover errfn", done => {
    errfn(res)("a phony message")
    done()
  })

  it("should have environment equals to testing", done => {
    environment.should.be.equal("testing")
    done()
  })
})