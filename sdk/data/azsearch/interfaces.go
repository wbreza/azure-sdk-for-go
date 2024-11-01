// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azsearch

// CharFilterClassification provides polymorphic access to related types.
// Call the interface's GetCharFilter() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CharFilter, *MappingCharFilter, *PatternReplaceCharFilter
type CharFilterClassification interface {
	// GetCharFilter returns the CharFilter content of the underlying type.
	GetCharFilter() *CharFilter
}

// CognitiveServicesAccountClassification provides polymorphic access to related types.
// Call the interface's GetCognitiveServicesAccount() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CognitiveServicesAccount, *CognitiveServicesAccountKey, *DefaultCognitiveServicesAccount
type CognitiveServicesAccountClassification interface {
	// GetCognitiveServicesAccount returns the CognitiveServicesAccount content of the underlying type.
	GetCognitiveServicesAccount() *CognitiveServicesAccount
}

// DataChangeDetectionPolicyClassification provides polymorphic access to related types.
// Call the interface's GetDataChangeDetectionPolicy() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DataChangeDetectionPolicy, *HighWaterMarkChangeDetectionPolicy, *SQLIntegratedChangeTrackingPolicy
type DataChangeDetectionPolicyClassification interface {
	// GetDataChangeDetectionPolicy returns the DataChangeDetectionPolicy content of the underlying type.
	GetDataChangeDetectionPolicy() *DataChangeDetectionPolicy
}

// DataDeletionDetectionPolicyClassification provides polymorphic access to related types.
// Call the interface's GetDataDeletionDetectionPolicy() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DataDeletionDetectionPolicy, *SoftDeleteColumnDeletionDetectionPolicy
type DataDeletionDetectionPolicyClassification interface {
	// GetDataDeletionDetectionPolicy returns the DataDeletionDetectionPolicy content of the underlying type.
	GetDataDeletionDetectionPolicy() *DataDeletionDetectionPolicy
}

// IndexerDataIdentityClassification provides polymorphic access to related types.
// Call the interface's GetIndexerDataIdentity() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *IndexerDataIdentity, *IndexerDataNoneIdentity, *IndexerDataUserAssignedIdentity
type IndexerDataIdentityClassification interface {
	// GetIndexerDataIdentity returns the IndexerDataIdentity content of the underlying type.
	GetIndexerDataIdentity() *IndexerDataIdentity
}

// IndexerSkillClassification provides polymorphic access to related types.
// Call the interface's GetIndexerSkill() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureOpenAIEmbeddingSkill, *ConditionalSkill, *CustomEntityLookupSkill, *DocumentExtractionSkill, *EntityLinkingSkill,
// - *EntityRecognitionSkill, *EntityRecognitionSkillV3, *ImageAnalysisSkill, *IndexerSkill, *KeyPhraseExtractionSkill, *LanguageDetectionSkill,
// - *MergeSkill, *OcrSkill, *PIIDetectionSkill, *SentimentSkill, *SentimentSkillV3, *ShaperSkill, *SplitSkill, *TextTranslationSkill,
// - *WebAPISkill
type IndexerSkillClassification interface {
	// GetIndexerSkill returns the IndexerSkill content of the underlying type.
	GetIndexerSkill() *IndexerSkill
}

// LexicalAnalyzerClassification provides polymorphic access to related types.
// Call the interface's GetLexicalAnalyzer() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CustomAnalyzer, *LexicalAnalyzer, *LuceneStandardAnalyzer, *PatternAnalyzer, *StopAnalyzer
type LexicalAnalyzerClassification interface {
	// GetLexicalAnalyzer returns the LexicalAnalyzer content of the underlying type.
	GetLexicalAnalyzer() *LexicalAnalyzer
}

// LexicalTokenizerClassification provides polymorphic access to related types.
// Call the interface's GetLexicalTokenizer() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ClassicTokenizer, *EdgeNGramTokenizer, *KeywordTokenizer, *KeywordTokenizerV2, *LexicalTokenizer, *LuceneStandardTokenizer,
// - *LuceneStandardTokenizerV2, *MicrosoftLanguageStemmingTokenizer, *MicrosoftLanguageTokenizer, *NGramTokenizer, *PathHierarchyTokenizerV2,
// - *PatternTokenizer, *UaxURLEmailTokenizer
type LexicalTokenizerClassification interface {
	// GetLexicalTokenizer returns the LexicalTokenizer content of the underlying type.
	GetLexicalTokenizer() *LexicalTokenizer
}

// ScoringFunctionClassification provides polymorphic access to related types.
// Call the interface's GetScoringFunction() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DistanceScoringFunction, *FreshnessScoringFunction, *MagnitudeScoringFunction, *ScoringFunction, *TagScoringFunction
type ScoringFunctionClassification interface {
	// GetScoringFunction returns the ScoringFunction content of the underlying type.
	GetScoringFunction() *ScoringFunction
}

// SimilarityClassification provides polymorphic access to related types.
// Call the interface's GetSimilarity() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BM25Similarity, *ClassicSimilarity, *Similarity
type SimilarityClassification interface {
	// GetSimilarity returns the Similarity content of the underlying type.
	GetSimilarity() *Similarity
}

// TokenFilterClassification provides polymorphic access to related types.
// Call the interface's GetTokenFilter() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ASCIIFoldingTokenFilter, *CjkBigramTokenFilter, *CommonGramTokenFilter, *DictionaryDecompounderTokenFilter, *EdgeNGramTokenFilter,
// - *EdgeNGramTokenFilterV2, *ElisionTokenFilter, *KeepTokenFilter, *KeywordMarkerTokenFilter, *LengthTokenFilter, *LimitTokenFilter,
// - *NGramTokenFilter, *NGramTokenFilterV2, *PatternCaptureTokenFilter, *PatternReplaceTokenFilter, *PhoneticTokenFilter,
// - *ShingleTokenFilter, *SnowballTokenFilter, *StemmerOverrideTokenFilter, *StemmerTokenFilter, *StopwordsTokenFilter, *SynonymTokenFilter,
// - *TokenFilter, *TruncateTokenFilter, *UniqueTokenFilter, *WordDelimiterTokenFilter
type TokenFilterClassification interface {
	// GetTokenFilter returns the TokenFilter content of the underlying type.
	GetTokenFilter() *TokenFilter
}

// VectorSearchAlgorithmConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetVectorSearchAlgorithmConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ExhaustiveKnnAlgorithmConfiguration, *HnswAlgorithmConfiguration, *VectorSearchAlgorithmConfiguration
type VectorSearchAlgorithmConfigurationClassification interface {
	// GetVectorSearchAlgorithmConfiguration returns the VectorSearchAlgorithmConfiguration content of the underlying type.
	GetVectorSearchAlgorithmConfiguration() *VectorSearchAlgorithmConfiguration
}

// VectorSearchCompressionConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetVectorSearchCompressionConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BinaryQuantizationCompressionConfiguration, *ScalarQuantizationCompressionConfiguration, *VectorSearchCompressionConfiguration
type VectorSearchCompressionConfigurationClassification interface {
	// GetVectorSearchCompressionConfiguration returns the VectorSearchCompressionConfiguration content of the underlying type.
	GetVectorSearchCompressionConfiguration() *VectorSearchCompressionConfiguration
}

// VectorSearchVectorizerClassification provides polymorphic access to related types.
// Call the interface's GetVectorSearchVectorizer() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureOpenAIVectorizer, *VectorSearchVectorizer, *WebAPIVectorizer
type VectorSearchVectorizerClassification interface {
	// GetVectorSearchVectorizer returns the VectorSearchVectorizer content of the underlying type.
	GetVectorSearchVectorizer() *VectorSearchVectorizer
}

