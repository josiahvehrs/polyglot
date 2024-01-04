import cli from "command-line-args";

export type ProjectorOpts = {
    args?: string[];
    pwd?: string[];
    config?: string;
};

export default function getOpts(): ProjectorOpts {
    return cli([
        {
            name: "arguments",
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
    ]) as ProjectorOpts;
}
