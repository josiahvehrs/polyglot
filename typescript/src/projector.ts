import { Config } from "./config";
import fs from "node:fs";
import path from "node:path";

export type Data = {
    projector: {
        // pwd
        [key: string]: {
            // key -> value
            [key: string]: string;
        };
    };
};

const defaltData = {
    projector: {},
};

export default class Projector {
    constructor(private config: Config, private data: Data) {}

    getValueAll(): { [key: string]: string } {
        let curr = this.config.pwd;
        let prev = "";
        const paths = [];

        do {
            prev = curr;
            paths.push(curr);
            curr = path.dirname(curr);
        } while (curr !== prev);

        return paths.reverse().reduce((acc, path) => {
            const value = this.data.projector[path];
            if (value) {
                Object.assign(acc, value);
            }
            return acc;
        }, {});
    }

    getValue(key: string): string | undefined {
        let curr = this.config.pwd;
        let prev = "";

        let out: string | undefined = undefined;

        do {
            const value = this.data.projector[curr]?.[key];
            if (value) {
                out = value;
                break;
            }

            prev = curr;
            curr = path.dirname(curr);
        } while (curr !== prev);

        return out;
    }

    setValue(key: string, value: string) {
        let dir = this.data.projector[this.config.pwd];
        if (!dir) {
            dir = this.data.projector[this.config.pwd] = {};
        }

        dir[key] = value;
    }

    removeValue(key: string) {
        const dir = this.data.projector[this.config.pwd];
        if (dir) {
            delete dir[key];
        }
    }

    save() {
        const configPath = path.dirname(this.config.config);
        if (!fs.existsSync(configPath)) {
            fs.mkdirSync(configPath, {
                recursive: true,
            });
        }
        fs.writeFileSync(this.config.config, JSON.stringify(this.data));
    }

    static fromConfig(config: Config): Projector {
        let data: Data;
        try {
            data = JSON.parse(fs.readFileSync(config.config).toString());
        } catch (error) {
            data = defaltData;
        }
        return new Projector(config, data);
    }
}
