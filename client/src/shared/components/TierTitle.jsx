import React from 'react';

export const details = ({
  'gold-bespoke': {
    text: "Gold Bespoke",
    icon: "/assets/images/icons/gold_bespoke.png",
  },
  gold: {
    text: "Gold",
    icon: "/assets/images/icons/gold.png",
  },
  'silver-bespoke': {
    text: "Silver Bespoke",
    icon: "/assets/images/icons/silver_bespoke.png",
  },
  silver: {
    text: "Silver",
    icon: "/assets/images/icons/silver.png",
  },
  bronze: {
    text: "Bronze",
    icon: "/assets/images/icons/bronze.png",
  },
});

export default class TierTitle extends React.Component {


  render() {
    const { tier } = this.props;
    const { icon, text } = details[tier];

    return (
      <div className="tier-title">
        <img src={icon} />
        <h3 className="inline">{text}</h3>
      </div>
    );
  }
}
