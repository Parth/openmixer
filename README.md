# Open Mixer: [openmixer.io](openmixer.io)

Note to reviewer: the more I explored this space, the more convinced I became that there was an opportunity to actually launch a new mixer. My motivation is outlined in the first section, and I'd love to get feedback about the validity of my thinking from other engineers (which is why I'm interested in working at gemini). However for the busy engineer that's just looking to evaluate my implementation of the Jobcoin mixer, I discuss it [here](#implementation), and have published my solution at [openmixer.io](openmixer.io).

## Exploration of crypto mixers

Open Mixer is a POC for a digital asset mixing service which focuses on transparency. This POC mixes [JobCoins](https://jobcoin.gemini.com/sanitary/api) - a "dummy" digital asset that's significantly easier to work with. 

As most blockchains are transparent ledgers, a mixers primary utility is provided by moving transaction information "off-chain" (generally to a service's servers). For example, without using a mixing service, once your barista recieves your coffee transaction they would be able to view what the input of that tx was (and the input of that one, and the input of that one, etc). Ultimately he or she would have a pretty good chance at being able to guess your income, and your employer would be able to tell what you're doing with your salary. Not ideal.

Mixing services are one of the many ways you can increase your privacy when using public ledgers. The idea is that when your barista goes to lookup the inputs of your transaction they go to a transaction with many inputs in them, and lots of outputs in them. When many people use a mixer, who's sending what to whom is not stored on the blockchain, but rather on the mixer itself.

Unfortunately this requires that you trust the mixing service with your funds, though there are some trustless ways to perform mixing ([Coinjoin](https://en.bitcoin.it/wiki/CoinJoin) for example), though these services require specialized wallets.

While exploring the space of mixers I had trouble finding an open source mixer which is available without Tor. Furthermore I could not find any mixing service which made it's current transactional volume public, a critical piece of information that you'd want prior to sending your transaction through the mixer. Open Mixer aims to make transparency it's highest priority. 

## Implementation 


