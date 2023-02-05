import chai from "chai"
import chaiHttp from "chai-http"

import {app} from "./index.mjs"
import {database, doMigrate} from "./configs/database.mjs";

chai.should();
chai.use(chaiHttp);

describe("Base API test", () => {

    before(async () => await doMigrate())
    after(async () => await database.destroy())

    it("Should be in testing mode", done => {
        if (!process.env.NODE_ENV) return done(new Error("NODE_ENV vazio"));
        process.env.NODE_ENV.should.be.eql("test");
        done();
    })

    it("Should be 'online'", done => {
        chai
            .request(app.callback())
            .get('/status')
            .end((err, res) => {
                if (err) return done(err);
                res.should.have.status(200);
                res.text.should.be.eql("online");
                done();
            });
    });

    it("Should list beers", done => {
        chai
            .request(app.callback())
            .get('/beers')
            .end((err, res) => {
                if (err) return done(err);
                res.should.have.status(200);
                res.body.should.be.an('Array');
                done();
            });
    });

    it("Should find one beer", done => {
        chai
            .request(app.callback())
            .get('/beers/1')
            .end((err, res) => {
                if (err) return done(err);
                res.should.have.status(200);
                res.body.should.be.an('Object');
                done();
            });
    });

    it("Should insert one beer", done => {
        chai
            .request(app.callback())
            .post('/beers')
            .send({titlebeer: "Spaten", descriptionbeer:"A cerveja que criou o estilo Munich Helles"})
            .end((err, res) => {
                if (err) return done(err);
                res.should.have.status(200);
                res.body.should.be.an('Array');
                done();
            });
    });

    it("Should update one beer", done => {
        chai
            .request(app.callback())
            .put('/beers/1')
            .send({titlebeer: "Crystal", descriptionbeer:"Topa curtir com a gente?"})
            .end((err, res) => {
                if (err) return done(err);
                res.should.have.status(200);
                res.body.should.be.an('number');
                done();
            });
    });

    it("Should remove one beer", done => {
        chai
            .request(app.callback())
            .del('/beers/6')
            .end((err, res) => {
                if (err) return done(err);
                res.should.have.status(200);
                res.body.should.be.an('number');
                done();
            });
    });
})