use anyhow::Result;
use clap::Parser;
use polyglot::{config::Config, opts};

fn main() -> Result<()> {
    let opts: Config = opts::Opts::parse().try_into()?;
    println!("{:?}", opts);

    Ok(())
}
