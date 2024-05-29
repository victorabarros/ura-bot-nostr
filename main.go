package main

import (
	"context"
	"fmt"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

// ContextKey is a type for context keys
type ContextKey string

const urlKey ContextKey = "url"

var (
	WORKING_RELAYS = []string{
		"wss://adre.su",
		"wss://nos.lol",
		"wss://nostr.bitcoiner.social",
		"wss://offchain.pub",
		"wss://relay.damus.io",
		"wss://relay.momostr.pink",
		"wss://relay.mostr.pub",
	}

	OTHER_RELAYS = []string{
		"wss://adre.su",
		"wss://at.nostrworks.com",
		"wss://br.purplerelay.com",
		"wss://brb.io",
		"wss://btc.klendazu.com",
		"wss://deschooling.us",
		"wss://knostr.neutrine.com",
		"wss://lightningrelay.com",
		"wss://nfnitloop.com",
		"wss://nos.lol",
		"wss://nostr-pub.semisol.dev",
		"wss://nostr-pub.wellorder.net",
		"wss://nostr-relay.bitcoin.ninja",
		"wss://nostr-relay.derekross.me",
		"wss://nostr-relay.schnitzel.world",
		"wss://nostr-relay.texashedge.xyz",
		"wss://nostr-verified.wellorder.net",
		"wss://nostr.8777.ch",
		"wss://nostr.bch.ninja",
		"wss://nostr.bitcoiner.social",
		"wss://nostr.cercatrova.me",
		"wss://nostr.corebreach.com",
		"wss://nostr.dvdt.dev",
		"wss://nostr.easydns.ca",
		"wss://nostr.einundzwanzig.space",
		"wss://nostr.fmt.wiz.biz",
		"wss://nostr.fractalized.net",
		"wss://nostr.fractalized.ovh",
		"wss://nostr.massmux.com",
		"wss://nostr.maximacitadel.org",
		"wss://nostr.mom",
		"wss://nostr.mustardnodes.com",
		"wss://nostr.mwmdev.com",
		"wss://nostr.nodeofsven.com",
		"wss://nostr.noones.com",
		"wss://nostr.olwe.link",
		"wss://nostr.orba.ca",
		"wss://nostr.oxtr.dev",
		"wss://nostr.pobblelabs.org",
		"wss://nostr.roundrockbitcoiners.com",
		"wss://nostr.screaminglife.io",
		"wss://nostr.sectiontwo.org",
		"wss://nostr.semisol.dev",
		"wss://nostr.slothy.win",
		"wss://nostr.swiss-enigma.ch",
		"wss://nostr.thesamecat.io",
		"wss://nostr.vulpem.com",
		"wss://nostr.wine",
		"wss://nostr.xpedite-tech.com",
		"wss://nostr21.com",
		"wss://nostrrelay.com",
		"wss://offchain.pub",
		"wss://paid.no.str.cr",
		"wss://relay.current.fyi",
		"wss://relay.damus.io",
		"wss://relay.guggero.org",
		"wss://relay.hodl.ar",
		"wss://relay.magiccity.live",
		"wss://relay.momostr.pink",
		"wss://relay.mostr.pub",
		"wss://relay.rebelbase.site",
		"wss://relay.sovereign-stack.org",
		"wss://relay.stoner.com",
		"wss://relay.toastr.space",
		"wss://relayable.org",
		"wss://slick.mjex.me",
		"wss://strfry.chatbett.de",
		"wss://welcome.nostr.wine",
	}
)

func main() {
	ev := nostr.Event{
		CreatedAt: nostr.Now(),
		Kind:      nostr.KindTextNote,
		Content:   "content from go-nostr 99 example from relay",
	}

	var nsec string = ""
	_, s, e := nip19.Decode(nsec)
	if e != nil {
		panic(e)
	}
	sk := s.(string)
	if err := ev.Sign(sk); err != nil {
		fmt.Println("fail to sign event", err.Error())
	}

	for _, url := range WORKING_RELAYS {
		ctx := context.WithValue(context.Background(), urlKey, url)
		relay, err := nostr.RelayConnect(ctx, url)
		if err != nil {
			fmt.Println("fail to connect to relay", err.Error())
			continue
		}

		fmt.Println("posting event to: ", url)
		if err := relay.Publish(ctx, ev); err != nil {
			fmt.Println("fail to publish to relay", err.Error())
			continue
		}
	}

}
