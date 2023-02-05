import {app} from "./app/index.mjs"
import {database, doMigrate} from "./app/configs/database.mjs"

console.log("Iniciando banco e serviço");

doMigrate().then(async () => {
    if (!process.env.NODE_ENV) {
        console.log("informe o NODE_ENV");
        await database.destroy();
    } else {
        const modo = process.env.NODE_ENV;
        const porta = process.env.PORT;
        console.log(`modo [${modo}] porta [${porta}]`)
        app.listen(porta);
    }
});