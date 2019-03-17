# Open Mixer

## Exploration of crypto mixers

Open Mixer is a POC for a digital asset mixing service. This POC mixes [JobCoins](https://jobcoin.gemini.com/sanitary/api) - a "dummy" digital asset that's significantly easier to work with. 

As most blockchains are transparent ledgers, a mixers primary utility is provided by moving transaction information "off-chain" (generally to a service's servers). For example, without using a mixing service, once your barista recieves your coffee transaction they would be able to view what the input of that tx was (and the input of that one, and the input of that one, etc). Ultimately he or she would have a pretty good chance at being able to guess your income, and your employer would be able to tell what you're doing with your salary. Not ideal.

Mixing services are one of the many ways you can increase your privacy when using public ledgers. The idea is that when your barista goes to lookup the inputs of your transaction they go to a transaction with many inputs in them, and lots of outputs in them. When many people use a mixer, who's sending what to whom is not stored on the blockchain, but rather on the mixer itself.

Unfortunately this requires that you trust the mixing service with your funds, though there are some trustless ways to perform mixing ([Coinjoin](https://en.bitcoin.it/wiki/CoinJoin) for example), though these services require specialized wallets.
