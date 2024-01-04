import cli from "command-line-args";

export type Opts = {
    args?: string[];
    config?: string;
    pwd?: string;
};

export default function getOpts(): Opts {
    return cli([
        {
            name: "args",
            type: String,
            defaultOption: true,
            multiple: true,
        },
        {
            name: "config",
            alias: "c",
            type: String,
        },
        {
            name: "pwd",
            alias: "p",
            type: String,
        },
    ]) as Opts;
}
