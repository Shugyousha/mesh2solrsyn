mesh2solrsyn
============

This program is supposed to parse [MeSH](https://www.nlm.nih.gov/mesh/)
into a Solr synonym file that maps each term in MeSH to all synonyms of
its hyponyms (recursively).

The goal is to make Solr useful for cases where you want to search for
all more specific instances of a more general term . For example, ideally
you should be able to search for "Cancer" which then expands the query
to all synonyms of "Cancer" in MeSH and all synonyms of all its hyponyms
and, in turn, their hyponyms and so on (through the use of a SynonymFilter
in Solr).
 
Testing of this program has been superficial at best so use with care.

