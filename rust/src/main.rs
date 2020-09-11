use clap::{load_yaml, App};
use log::info;

fn flashy<T>(min: T, max: T)
where
    T: Into<f32>,
{
    use rand::distributions::Standard;
    use rand::prelude::*;
    use std::{thread, time};

    let min = min.into();
    let max = max.into();

    info!(
        "Initialised, now going to start doing my thing. Min: {:?}, Max: {:?}",
        min, max
    );

    loop {
        let min_rand: f32 = StdRng::from_entropy().sample(Standard);
        let max_rand: f32 = StdRng::from_entropy().sample(Standard);
        let min_sleep =
            time::Duration::from_millis(((1000_f32 * min) + (500_f32 * min_rand)) as u64);
        let max_sleep =
            time::Duration::from_millis(((max - min) * 1000_f32 + (100_f32 * max_rand)) as u64);
        thread::sleep(min_sleep);
        info!("Just slept for {:?} (min_sleep)", min_sleep);
        autopilot::mouse::click(autopilot::mouse::Button::Left, Some(0));
        thread::sleep(max_sleep);
        info!("Just slept for {:?} (max_sleep)", max_sleep);
        autopilot::mouse::click(autopilot::mouse::Button::Left, Some(0));
    }
}

fn main() {
    // The YAML file is found relative to the current file, similar to how modules are found
    let yaml = load_yaml!("cli.yml");
    let matches = App::from(yaml).get_matches();
    let min: f32 = matches
        .value_of("min")
        .unwrap_or("2")
        .to_string()
        .parse()
        .unwrap_or(2.0);
    let max: f32 = matches
        .value_of("max")
        .unwrap_or("3")
        .to_string()
        .parse()
        .unwrap_or(3.0);
    flashy(min, max);
}
