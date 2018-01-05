
const chai = require('chai');
const chaiHttp = require('chai-http');
const server = require('../src/main');

chai.should();
chai.use(chaiHttp);

describe("Beers test suite", _ => {

  const req = chai.request(server.app);

  let idbeer = 0

  it("should create one beer", done => {

    let beer = {
      titlebeer: "Itaipava",
      descriptionbeer: "A cerveja 100%",
    }

    req.post("/beer/save").send(beer).then(ret => {
      idbeer = ret.body.idbeer
      ret.status.should.be.equal(200)
      done()
    }).catch(err => done(err))

  })

  it("should delete one beer", done => {
    req.del(`/beer/${idbeer}`).then(ret => {
      ret.status.should.be.equal(200)
    }).catch(err => done(err))
  })
})