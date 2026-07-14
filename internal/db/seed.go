package db

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"strconv"

	"github.com/JayTailor45/go-social/internal/store"
)

var usernames = []string{
	"NovaFox",
	"ZenByte",
	"LunaRay",
	"PixelAsh",
	"EchoZen",
	"SkyVolt",
	"FrostIQ",
	"DriftFox",
	"NeonJet",
	"SolarHex",
	"VoidSnap",
	"BlazeZen",
	"RapidNova",
	"CyberAsh",
	"OrbitJay",
	"TurboLeaf",
	"GhostPixel",
	"HyperLynx",
	"FlareBit",
	"SilentHex",
	"StormJay",
	"CrystalFox",
	"ShadowOrb",
	"SwiftNova",
	"QuantumBee",
	"BlueDrift",
	"TurboNova",
	"PixelMint",
	"DarkPulse",
	"ZenSpark",
	"OrbitRush",
	"SilverByte",
	"FlashEcho",
	"AlphaMint",
	"VortexRay",
	"CloudHex",
	"WildNova",
	"FrostByte",
	"SkyBlink",
	"RapidGlow",
	"CyberLoom",
	"PulseFox",
	"DreamVolt",
	"NovaMint",
	"EchoBlaze",
	"TurboWave",
	"PixelNova",
	"ShadowMint",
	"VoltRush",
	"ZenOrbit",
	"FireLynx",
	"AstroMint",
	"BlitzFox",
	"CodeDrift",
	"NightVolt",
	"WavePixel",
	"SteelNova",
	"ArcticZen",
	"FlameByte",
	"GlowRush",
	"TurboZen",
	"PulseRay",
	"SilentNova",
	"OrbitGlow",
	"HexLoom",
	"RapidFox",
	"SkyMint",
	"StormByte",
	"CyberWave",
	"FrostNova",
	"PixelRush",
	"QuantumFox",
	"NovaDrift",
	"BlazeMint",
	"EchoRush",
	"DarkNova",
	"TurboAsh",
	"ZenLynx",
	"VoltPixel",
	"FlashNova",
	"DreamByte",
	"SolarRush",
	"CodeNova",
	"HyperMint",
	"WildByte",
	"CrystalZen",
	"OrbitLynx",
	"SwiftGlow",
	"NightFox",
	"NovaSpark",
	"PixelVolt",
	"BlueNova",
	"GhostRush",
	"StormMint",
	"CyberNova",
	"ArcticByte",
	"TurboPulse",
	"FlareNova",
	"ZenDrift",
	"ShadowWave",
}

var titles = []string{
	"Silent Rivers at Dawn",
	"Broken Compass Theory",
	"Neon Shadows Rising",
	"Forgotten Kingdoms Return",
	"Echoes Beyond Mountains",
	"Falling Stars Protocol",
	"Crimson Horizon Drift",
	"Midnight Lantern Society",
	"Shattered Crown Legacy",
	"Whispers Through Steel",
	"Frozen Ember Chronicles",
	"Golden Skies Collapse",
	"Hidden Truths Unlocked",
	"Stormwatch Over Eden",
	"Vanishing Roads Ahead",
	"Ancient Machines Awaken",
	"Burning Tides Approach",
	"Digital Nomads Revolt",
	"Last Voyage Home",
	"Velvet Thunder Awakens",
}

var contents = []string{
	"As the first light touched the valley, the rivers moved silently through the sleeping forests. Travelers believed these waters carried forgotten memories from ancient times.",
	"The old navigator claimed that broken compasses revealed a person's true destination instead of pointing north. Few explorers believed him until strange journeys started proving his theory correct.",
	"The city never truly slept beneath the endless glow of neon lights. Hidden within crowded streets, a resistance movement slowly prepared to challenge the system controlling everyone.",
	"Ancient banners once buried beneath desert sands began appearing across distant lands. Historians dismissed the signs until abandoned fortresses started glowing at night.",
	"For decades, strange echoes rolled through the mountains after sunset. Climbers described hearing voices calling their names from deep within the cliffs.",
	"When the first star fell from the sky, scientists realized it was not a meteor but a machine carrying unknown technology. Governments quickly activated a secret emergency protocol.",
	"Sailors feared the crimson horizon that appeared once every generation. Ships crossing into the red mist vanished without leaving any trace behind.",
	"Every midnight, mysterious lanterns appeared across the old city streets. Those who followed them discovered hidden gatherings protecting dangerous secrets.",
	"The shattered crown was believed lost forever after the empire collapsed. Years later, rumors of its return sparked conflicts across neighboring kingdoms.",
	"Steel factories echoed with whispers no worker could explain. Some believed the machines themselves had started communicating after decades of silence.",
	"Deep beneath frozen ruins, explorers uncovered embers that still burned after thousands of years. The discovery challenged everything scientists understood about energy.",
	"The golden skies appeared moments before entire cities disappeared without warning. Survivors spoke of massive shadows moving above the clouds.",
	"A hidden archive beneath the capital contained documents powerful enough to destroy governments. Only a small group knew how to unlock its secrets.",
	"Stormwatch towers stood across the coastline, warning citizens of unnatural storms approaching from distant oceans. No one understood what caused them.",
	"Travelers crossing the abandoned highways often reported roads vanishing behind them. Some claimed the paths were alive and constantly changing direction.",
	"Ancient machines buried underground suddenly awakened after centuries of silence. Engineers struggled to understand the strange symbols appearing across their surfaces.",
	"Massive tides began burning instead of flowing with water, terrifying coastal villages everywhere. Scientists rushed to investigate the impossible phenomenon.",
	"Digital nomads around the world secretly united against corporations controlling online communication. Their rebellion spread faster than governments could contain it.",
	"A damaged ship carrying the last survivors drifted endlessly through empty space. The crew searched desperately for a place they could finally call home.",
	"Thunderstorms unlike anything recorded before started appearing across remote regions. Witnesses described hearing voices hidden within the roaring skies.",
}

var tags = []string{
	"nature",
	"mystery",
	"river",
	"ancient",
	"adventure",
	"compass",
	"exploration",
	"journey",
	"fiction",
	"cyberpunk",
	"city",
	"resistance",
	"future",
	"technology",
	"kingdom",
	"history",
	"fantasy",
	"empire",
	"mountains",
	"echoes",
	"thriller",
	"supernatural",
	"space",
	"science",
	"aliens",
	"protocol",
	"ocean",
	"sailors",
	"secret",
	"society",
	"lanterns",
	"underground",
	"crown",
	"war",
	"legacy",
	"steel",
	"factory",
	"machines",
	"industrial",
	"frozen",
	"ruins",
	"discovery",
	"energy",
	"golden",
	"sky",
	"disaster",
	"survival",
	"archive",
	"government",
	"hidden",
	"secrets",
	"storm",
	"coastline",
	"weather",
	"warning",
	"roads",
	"travel",
	"awakening",
	"scifi",
	"tides",
	"fire",
	"digital",
	"rebellion",
	"freedom",
	"spaceship",
	"hope",
	"thunderstorm",
	"voices",
}

var comments = []string{
	"Absolutely loved this post!",
	"This was really interesting to read.",
	"Great perspective on the topic.",
	"I never thought about it this way.",
	"Very well written and engaging.",
	"This deserves more attention.",
	"Awesome content as always.",
	"I learned something new today.",
	"Really cool idea!",
	"This was surprisingly inspiring.",
	"Thanks for sharing this.",
	"Now this is quality content.",
	"I completely agree with this.",
	"Such an underrated topic.",
	"Looking forward to more posts like this.",
	"This kept me hooked till the end.",
	"Very creative and unique.",
	"Honestly, this was amazing.",
	"Simple but very impactful.",
	"This gave me a new perspective.",
	"Brilliant explanation.",
	"Your storytelling is excellent.",
	"This was worth the read.",
	"One of the best posts today.",
	"Really thought provoking.",
	"I enjoyed every bit of this.",
	"Super engaging content.",
	"This felt cinematic somehow.",
	"Interesting concept for sure.",
	"Nicely explained and easy to follow.",
	"This post has great energy.",
	"I would love a part two.",
	"Very immersive writing style.",
	"This deserves to go viral.",
	"Such a fascinating read.",
	"I can totally imagine this happening.",
	"Great balance of detail and simplicity.",
	"This instantly caught my attention.",
	"Very original content idea.",
	"I appreciate the effort behind this.",
	"This was unexpectedly deep.",
	"Really fun to read through.",
	"Your creativity stands out here.",
	"This has serious potential.",
	"Clean and concise writing.",
	"This made my day better.",
	"I would definitely recommend this.",
	"Strong opening and ending.",
	"This was both entertaining and informative.",
	"Can't wait to read more from you.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	tx, _ := db.BeginTx(ctx, nil)

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("error seeding user ", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error seeding post ", err)
			return
		}
	}

	comments := generateComments(500, posts, users)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, *comment); err != nil {
			log.Println("error seeding comment ", err)
			return
		}
	}

	log.Println("database seeding has been completed")
}

func generateUsers(count int) []*store.User {
	users := make([]*store.User, count)

	for i := 0; i < count; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + strconv.Itoa(i+1),
			Email:    usernames[i%len(usernames)] + strconv.Itoa(i+1) + "@test.com",
		}
	}
	return users
}

func generatePosts(count int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, count)

	for i := 0; i < count; i++ {
		posts[i] = &store.Post{
			UserID:  users[rand.Intn(len(users))].ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(count int, posts []*store.Post, users []*store.User) []*store.Comment {
	cmts := make([]*store.Comment, count)

	for i := 0; i < count; i++ {
		cmts[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cmts
}
