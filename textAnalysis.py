from itertools import zip_longest


def textPreprocessing(fileName):
	with open("Cleaned_" + fileName, "w", encoding='UTF8' ) as w:
		with open(fileName, "r", encoding='UTF8' ) as r:
			for line in r:
				wordDef = line.split('\t')
				# Removing chars which are used to denote ittonation etc as they will only cause confusion
				# When trying to match letters to ipa
				wordDef[1] = wordDef[1].replace('ˈ', "").replace("/", "").replace(":", "").replace("‍", "").replace("ː", "").replace("ˌ", "")
				w.write("\t".join(wordDef))


def charToIPAMatching(word, ipa):
	if len(word) == len(ipa):
		print(word)
		print(list(zip_longest(word, ipa)))


file = "en_UK.tsv"
textPreprocessing(file)

cleanedFile = "Cleaned_" + file

with open(cleanedFile, "r", encoding='UTF8' ) as r:
	for line in r:
		wordDef = line.replace("\n", "").split("\t")
		charToIPAMatching(wordDef[0], wordDef[1])