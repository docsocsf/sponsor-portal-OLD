import React from 'react';
import Sponsor from './components/Sponsor';
import TierTitle from 'Components/TierTitle';
import { fetchSponsors } from '../actions';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';

export default class SponsorPanel extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
      sponsors: [],
      activeIndex: 0,
    }
  }

  async componentDidMount() {
    let sponsors = await fetchSponsors();
    this.setState({ sponsors });
  }

  handleClick = (e, titleProps) => {
    const { index } = titleProps;
    const { activeIndex } = this.state;
    const newIndex = activeIndex === index ? -1 : index;

    this.setState({ activeIndex: newIndex });
  }

  render() {
    const { activeIndex, sponsors } = this.state;
    const titles = sponsors.map((t, i) => (
      <Tab key={i}>
        <TierTitle tier={t.tier} />
      </Tab>
    ));

    const tiers = sponsors.map((t, i) => {
      let sponsors = t.sponsors.map((s, i) => (
        <Sponsor key={i} info={s} tier={t.tier} />
      ));

      return (
        <TabPanel key={i}>
          { sponsors }
        </TabPanel>
      );
    });

    return (
      <section id="sponsors">
        <h2>Sponsors</h2>
        <Tabs className="vertical tabs">
          <TabList>
            {titles}
          </TabList>
          {tiers}
        </Tabs>
      </section>
    );
  }
}
