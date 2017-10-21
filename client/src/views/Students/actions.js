import fetchWithConfig from '../../fetch';
import { getJWTHeader } from '../../jwt';

export const fetchUser = async () => {
  const headers = await getJWTHeader();
  let resp = await fetchWithConfig('/api/students/user', { headers })
  return resp.body
}

export const fetchCV = async () => {
  const headers = await getJWTHeader();
  let resp = await fetchWithConfig('/api/students/cv', { headers })
  return resp.body
}


export const fetchSponsors = async () => {
  return Promise.resolve(sponsors);
}

const sponsors = [
  {
    "tier": "gold",
    "sponsors": [
      {
        "name": "Facebook",
        "url": "https://www.facebook.com/",
        "logo": "facebook_logo.png",
        "bespoke": true,
        "description": 'The chance to move fast, be bold and build products and services with impact has never been greater. At Facebook, we have a saying the journey is only 1% finished — join us as you begin yours.\nWe have tons of exciting opportunities available for university grads and interns who want to help us in our mission to bring the world closer together.'
      },
      {
        "name": "Palantir",
        "url": "https://www.palantir.com/",
        "email": "imperial@palantir.com",
        "logo": "palantir_logo.png",
        "bespoke": true,
        "description": "Palantir builds the most sophisticated data integration and analysis so ware in the world. We partner with government agencies, commercial institutions, and non-profit organizations to transform how they use data and technology.\nCome work with us. Work with the best and brightest minds to solve urgent problems of massive scale. Work across an ever expanding set of domains and industries. Work on the problems you read about on the front page of the newspaper.",
      },
      {
        "name": "NewVoiceMedia",
        "url": "https://www.newvoicemedia.com/",
        "apply": "http://www.newvoicemedia. com/ about-newvoicemedia/careers/",
        "logo": "newvoicemedia_logo.png",
        "description": "NewVoiceMedia is in the process of disrupting one of the biggest markets in the world - business telephony. Well- funded by cloud-focused VC’s, and with a management team with a decade of cloud experience, we have a strong vision for the future and aggressive global growth plans. We are looking for exciting, passionate, innovative people to join our fast growing team and help us achieve our goals.",
      },
      {
        "name": "G-Research",
        "url": "https://www.gresearch.co.uk",
        "logo": "g_research_logo.png",
        "description": "G-Research develops and maintains an incredibly successful set of research, analysis and real-time platforms used within the quantitative finance arena.\nWe are currently in the middle of transforming these platforms from a suite of mostly in-house enterprise solutions to cloud style distributed systems that leverage the best of open source and third party offerings.\nJoin us and you will be at the centre of this exciting change and will have the opportunity to work with technologies such as Spark, Kafka, Cassandra, .NET Core, Docker, Kubernetes and more as we make the transition.\nAs we evolve our technology we are also evolving the way we work. Innovation happens at the blurry edge of disciplines so we work cross - functionally to generate new ideas and business opportunities.\nWe’re hiring the brightest Computer Science or JMC students to join us for 4 or 6 month placements where you will get the chance to fully immerse yourself in life at G-Research and all we have to offer. From day 1, you will be allocated a real project to work on alongside some of the best engineers in the world, with a dedicated mentor to guide you during your time with us.",
      },
      {
        "name": "Netcraft",
        "url": "https://www.netcraft.com",
        "email": "cv@netcraft.com",
        "logo": "netcraft_logo.png",
        "description": "A world-leading security services company, performing penetration testing and anti-fraud services to critical institutions throughout the world. Netcra  wants to recruit so ware developers to its team based in Bath, England.\nProjects involve working on a range of commercial services extending Netcra ’s network exploration and internet security services.\nNetcraft’s services protect some of the worlds most important institutions, and its phishing site feed is consumed by Google, Mozilla, Apple, Opera, and Microsoft.",
      },
      {
        "name": "NextJump",
        "url": "https://www.nextjump.com",
        "logo": "next_jump_logo.png",
        "description": "Next Jump, as a ‘Deliberately Developmental Organisation’, is on a mission to create high-performing workplace cultures across the globe. People development is at the heart of everything we do and we follow this simple mantra: Better Me + Better You = Better Us. Next Jumpers train and improve themselves at work (Better Me) and use that knowledge to help others (Better You), together we create a better world (Better Us).\nHeadquartered in New York, based in London, our cutting-edge e-commerce platform is used by 30 million employees across Fortune 1000 companies and generates $2 billion in sales every year. We use the revenue generated from this platform to fuel our social movement: changing workplace culture. We run ‘Leadership Academies’ and open source our technology to help organisations build adaptive learning teams. We have helped a large range of organisations, varying from likes of JP Morgan Chase, the Military through to state schools/ charities.\nAs our core and social businesses continue to grow rapidly, we are looking for driven and entrepreneurial individuals to join us! If you are looking for a ‘not- your-typical graduate job/internship’, then Next Jump’s the place for you!",
      },
      {
        "name": "Accenture",
        "url": "https://www.accenture.com/",
        "logo": "accenture_logo.png",
        "description": "Accenture solves our clients’ toughest challenges by providing unmatched services in strategy, consulting, digital, technology and operations. We partner with more than three-quarters of the Fortune Global 500, driving innovation to improve the way the world works and lives. With expertise across more than 40 industries and all business functions, we deliver transformational outcomes for a demanding new digital world.",
      },
      {
        "name": "GoCardless",
        "url": "https://gocardless.com",
        "logo": "go_cardless_logo.png",
        "description": "GoCardless is building the payments network for the internet. We’re bringing Direct Debit into the digital age, building the best system for taking recurring payments.\nGoCardless is creating a new international payments network to rival credit and debit cards. Our ambition is to break down barriers so businesses can quickly and easily take payments from anyone, anywhere in the world.\nThe calibre of people that you will be surrounded by at GoCardless is second to none. You’ll be in a fantastic team of smart, talented and intellectually curious people who are always looking to learn and support each other.\nWe use whichever methodologies help us be most e ective. Although we work in an agile way, each project team organises how they work, how they run their standups, and how they scope in a way that best  ts them. We deploy many times a day, supported by a test suite we trust and a culture of code review.",
      }
    ]
  },
  {
    "tier": "silver",
    "sponsors": [
      {
        "name": "Microsoft",
        "url": "https://www.microsoft.com",
        "logo": "microsoft_logo.png",
        "bespoke": true,
        "description": "No matter what your passion is, you’ll  nd it here. Imagine the opportunities you’ll have in a company with more than 100,000 employees in more than 100 countries, working on hundreds of products—spanning games, phones, developer tools, business solutions and operating systems. We work hard, but we value work/life balance, and each of us defines what that means to us. So why not explore what we do, where we do it, and what life is really like at Microsoft. You just might be surprised.",
      },
      {
        "name": "Bloomberg",
        "url": "https://www.bloomberg.com",
        "apply": "https://www.bloomberg.com/careers/technology/engineering/",
        "logo": "bloomberg_logo.png",
        "description": "Born in 1981, Bloomberg is the world’s primary distributor of financial data and a top news provider of the 21st century. A global information and technology company, we use our dynamic network of data, ideas and analysis to solve di icult problems. Our customers around the world rely on us to deliver accurate, real-time business and market information that helps them make important  nancial decisions.\nAt Bloomberg, we are guided by four core values that are the foundation of our continued success: innovation, collaboration, customer service and doing the right thing.\nWe harness the power of data and analytics to organize, understand and improve our world. It’s our purpose. Come find yours.\n",
      },
      {
        "name": "Entrepreneur First",
        "url": "https://www.joinef.com",
        "logo": "ef_logo.png",
        "description": "Entrepreneur First brings together extraordinary people to build startups from scratch.\nAmbitious individuals leave places like Google, Goldman Sachs and Stanford to join EF. Why? Because starting a startup is the highest impact thing you can do with your life.\nEF is the best place to find a co-founder, build a company and access the world’s best investors. We’ve helped build over 80 technology companies, collectively worth over $1bn.\nThe greatest leaps in technology and society have been driven by a handful of obsessive, dedicated and diverse individuals.\nEF ampli es your e orts from the very beginning of your journey. EF funds you as you build your cofounding team, develop your idea, and accelerate through fundraising from the world’s best investors.\nEF is backed by Reid Hoffman (founder of LinkedIn), Greylock Partners, Mosaic Ventures, Founders Fund, Lakestar Capital, and Deep Mind founders Demis Hassabis and Mustafa Suleyman.\nIf your ambition outweighs your current career status, apply to join us.\nEvery six months we bring together a selection of mission-driven individuals who want to build tech companies with a global impact. We spend hundreds of hours screening thousands of pro les and select up to 100 potential founders to join our community.\nWe don’t take a cookie-cutter approach to selecting founders;  rst and foremost we like to work with highly ambitious people who have strong and unconventional beliefs about what the future looks like.",
      },
      {
        "name": "Thread",
        "url": "https://www.thread.com",
        "logo": "thread_logo.png",
        "description": "Our goal is to make it easy to dress well. We believe if we get this right we can create the new default for how hundreds of millions of people across the world buy their clothes.\nThread brings people a shortlist of the perfect clothes just for them, personalised to their size, budget and style. We achieve this by fusing together talented human stylists with powerful machine learning algorithms.\nChanging the way the retail industry works is ambitious. One of our core strategies for getting there is to create one of the highest quality technology companies anywhere in the world. Our founders have always been passionate about recruiting, and are committed to building an enduring, long-term company  lled with the most talented people on the planet. We have a lot of work to do, and are looking for people who want to make an impact and help us build a global brand. As a company we are committed to building a strong, excellence-focused culture. We place high value on learning and personal growth. If you’re driven, care deeply about your cra  and want an unrivalled opportunity to develop your skill set in a fast-growing business, Thread might be for you.",
      }
    ]
  },
  {
    "tier": "bronze",
    "sponsors": [
      {
        "name": "Wluper",
        "url": "http://wluper.com",
        "email": "imperial@wluper.com",
        "logo": "wluper_logo.png",
        "description": "Voice and Chat as an interface is the next big step in making so ware accessible. Conversational AI interfaces will soon enable people to communicate naturally with machines. Wluper is developing a personal intelligent assistant for navigation and transportation, and the fundamental AI technology needed to allow for easy, intuitive journey planning and a new class of user experience.\nWe are working on deep learning NLP and AI technology to revolutionise mobile voice enabled agents for all transit-related conversations. We believe now is the right time to build this product and we have assembled a great team of very talented researchers and engineers in order to do it. Backed by Jaguar Land Rover’s InMotion Ventures, we are always looking for keen and skilled talents who will have signi cant in uence on the whole project-lifecycle, our overall strategy by helping de ne system features, and spearhead the best practices that enable a quality product. Come join us, together we will work on one of the coolest and most innovative  elds of technology in the recent years.",
      },
      {
        "name": "BlackRock",
        "url": "https://www.blackrock.com/",
        "apply": "https://careers.blackrock.com/",
        "logo": "blackrock_logo.png",
        "description": "BlackRock combines the best of being a technology pioneer and  nancial leader. Since our inception in 1988, we’ve grown from an entrepreneurial start-up into the world’s largest investment manager by putting technology - and continuous innovation - at the heart of our business.\nWe’ve invented our own global enterprise investment system, ‘Aladdin’: a technology ecosystem that empowers our colleagues to do better business every day. It is an unmatched operating platform - and central nervous system - that unites all the information, people and technology needed to manage money in real time and at every step in the investment process.\nUnlike other financial firms, so ware engineering at BlackRock is a profit centre, not a support function, and stands on its own as a fast-growing business. Developers work in teams focused on speci c components of Aladdin and use innovative new technologies to build solutions for BlackRock and many Aladdin clients.",
      },
      {
        "name": "J.P. Morgan",
        "url": "https://www.jpmorgan.com/",
        "apply": "https://careers.jpmorgan.com",
        "logo": "jpmorgan_logo.png",
        "description": "Your Career. Your Way.\nAt J.P. Morgan, we are committed to helping businesses, markets and communities grow and develop in more than 100 countries. Working with us means you’ll learn from our team of experts in a supportive and collaborative environment and gain the skills to make a direct contribution to a firm with a legacy lasting over 200 years.\nWe want to see your creativity, communications skills and drive. While your academic achievements are important, we’re also looking for your individuality and passion as demonstrated by extra-curricular activities. We want to help you ful l your potential as you build your career here.",
      },
      {
        "name": "SIG",
        "url": "https://sig.com",
        "apply": "https://www.sig.com/campus-programmes/",
        "logo": "sig_logo.png",
        "description": "Game theory rings true in everything we do. We’re big on competition, strategy, and managing risk – just like great gamers. We are problem solvers who analyse the financial markets to identify profitable trading opportunities.\nSIG is one of the largest privately-held  nancial institutions in the world. We have o ices all around the globe and trade an extensive range of  nancial products. Building virtually all of our own trading technology from scratch, we are a leading and innovative company in high-performance, low latency trading.\nOur trading internship is a 10-week summer programme, designed to expose you to the working environment that you can expect if you return to SIG as a full-time Assistant Trader. The internship programme is our primary hiring source for full-time trading positions at SIG.\nWe build our talent from the ground up. Students who start their career at SIG learn what it’s like to work in a trading environment through a combination of SIG’s classroom education and impactful work on our trading desks.\nUpon completion of the internship you may be offered a returning role as an Assistant Trader where you will become responsible for making independent trading decisions.",
      },
      {
        "name": "Bank of America Merrill Lynch",
        "url": "https://www.bofaml.com",
        "logo": "baml_logo.png"
      },
      {
        "name": "Capital One",
        "url": "https://www.capitalone.co.uk",
        "logo": "capital_one_logo.png",
        "description": "Capital One is one of the UK’s leading credit card companies and a Fortune 200 company. Still led by our founder and global CEO, Rich Fairbank, we are on a mission to bring ingenuity, simplicity and humanity to an industry that’s ripe for change. We focus on delivering products and services with our customers’ needs at their heart. We get closer to them through the acres of data their card use yields and analyse it so we know how and where we can help them succeed with credit. The result is a data-driven, tech-enabled business on a mission to simplify  nance and powered by the ideas, passion and commitment of our people.\nWorking in technology at Capital One isn’t just about the tech. It’s about the tech journey. Design thinking, open source so ware, public cloud, APIs, a startup mind-set and agile framework; these are the secrets to our success. Being more agile means we get more from technology. It means we’re better able to solve business challenges. You don’t necessarily need a technical degree as we believe technical understanding can be taught. What we look for are tech-obsessed problem solvers with inquisitive minds and curious natures who can create innovative solutions.",
      }
    ]
  }
];
