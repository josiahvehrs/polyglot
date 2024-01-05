use anyhow::Result;
use clap::Parser;
use polyglot::{
    config::{Config, Operation},
    opts,
    projector::Projector,
};

fn main() -> Result<()> {
    let config: Config = opts::Opts::parse().try_into()?;
    let mut proj = Projector::from_config(config.config, config.pwd);

    match config.operation {
        Operation::Print(None) => {
            let value = proj.get_value_all();
            let value = serde_json::to_string(&value)?;
            println!("{}", value);
        }
        Operation::Print(Some(k)) => {
            if let Some(v) = proj.get_value(&k) {
                println!("{}", v)
            }
        }
        Operation::Add(k, v) => {
            proj.set_value(k, v);
            proj.save()?;
        }
        Operation::Remove(k) => {
            proj.remove_value(&k);
            proj.save()?;
        }
    }

    Ok(())
}
