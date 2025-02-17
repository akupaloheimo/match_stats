# Matsi statsit (ohjeet)

Kirjoita Golang ohjelma, joka lukee lyöntidatan, laskee niiden perustella pelaajien ottelutilastot ja tallentaa tilastot json encodattuina [protobuffeina](https://protobuf.dev/).


1. Lyöntien lukeminen
Lyönnit löytyvät json encodattuina protobuffeina tiedostosta shots.json. Lyönti proto on määritelty shot.proto tiedostossa. shot.proto tiedoston avulla pystyt generoimaan go koodin jolla shots.json lukeminen on helppoa.

2. Ottelutilastojen laskeminen
Shotti datan perusteella pystyy laskemaan tyypillisiä tennisottelun tilastoja. Tilastoissa ei tarvitse olla kuin muutama tyypillinen tennisottelun tilasto. Toivottavasti shotti data on oikein, mutta sielä saatta tietysti olla joitain virheitä, koska tein sen käsin. Voit halutessasi lisätä shotteja.

3. Ottelutilastojen tallentaminen
Ottelutilasto proton voit suunnitella itse ja tallentaa json formaatissa. 

# How to run the program

Make sure you have the proper version of Golang (Go 1.23.5 should work)

in the match_stats directory you can run "go run shots.go" and the results should be seen in match_stats.json
