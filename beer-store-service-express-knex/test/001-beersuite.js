// cross-env NODE_ENV=testing knex migrate:latest ; cross-env NODE_ENV=testing mocha --timeout=30000 --exit
const chai = require('chai')
const chaiHttp = require('chai-http')
const server = require('../src/main')

chai.should()
chai.use(chaiHttp)

describe("Beers test suite", _ => {

  let req = chai.request(server.app)
  
  let beer = {
    titlebeer: "Itaipava",
    descriptionbeer: "A cerveja 100%",
  }
  
  let idbeer = 0

  it("should create one beer", done => {

    req.post("/beer/save").send(beer).end((err, ret) => {
      idbeer = ret.body.idbeer
      ret.status.should.be.equal(200)
      done(err)
    })
  })

  it("should list at least one beer", done => {
    req.get("/beer/list").end((err, ret) => {
      ret.body.should.have.lengthOf.above(1)
      done(err)
    })
  })

  it("should update beer description", done => {
    const descriptionbeer = "A cerveja que desce redondo"
    req.put("/beer/save").send({ idbeer, descriptionbeer }).end((err, ret) => {
      ret.body.descriptionbeer.should.be.equal(descriptionbeer)
      done(err)
    })
  })

  it("should delete one beer", done => {
    req.del(`/beer/${idbeer}`).end((err, ret) => {
      ret.status.should.be.equal(200)
      done(err)
    })
  })
})