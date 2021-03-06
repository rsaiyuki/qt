#ifdef __cplusplus
extern "C" {
#endif

void QAbstractMessageHandler_DestroyQAbstractMessageHandler(void* ptr);
void QAbstractUriResolver_DestroyQAbstractUriResolver(void* ptr);
int QAbstractXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2);
int QAbstractXmlNodeModel_Kind(void* ptr, void* ni);
char* QAbstractXmlNodeModel_StringValue(void* ptr, void* n);
void* QAbstractXmlNodeModel_TypedValue(void* ptr, void* node);
void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(void* ptr);
void QAbstractXmlReceiver_Attribute(void* ptr, void* name, void* value);
void QAbstractXmlReceiver_Characters(void* ptr, void* value);
void QAbstractXmlReceiver_Comment(void* ptr, char* value);
void QAbstractXmlReceiver_EndDocument(void* ptr);
void QAbstractXmlReceiver_EndElement(void* ptr);
void QAbstractXmlReceiver_EndOfSequence(void* ptr);
void QAbstractXmlReceiver_NamespaceBinding(void* ptr, void* name);
void QAbstractXmlReceiver_ProcessingInstruction(void* ptr, void* target, char* value);
void QAbstractXmlReceiver_StartDocument(void* ptr);
void QAbstractXmlReceiver_StartElement(void* ptr, void* name);
void QAbstractXmlReceiver_StartOfSequence(void* ptr);
void QAbstractXmlReceiver_DestroyQAbstractXmlReceiver(void* ptr);
char* QSimpleXmlNodeModel_StringValue(void* ptr, void* node);
void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(void* ptr);
void* QSourceLocation_NewQSourceLocation();
void* QSourceLocation_NewQSourceLocation2(void* other);
void* QSourceLocation_NewQSourceLocation3(void* u, int l, int c);
int QSourceLocation_IsNull(void* ptr);
void QSourceLocation_SetUri(void* ptr, void* newUri);
void QSourceLocation_DestroyQSourceLocation(void* ptr);
void* QXmlFormatter_NewQXmlFormatter(void* query, void* outputDevice);
void QXmlFormatter_Attribute(void* ptr, void* name, void* value);
void QXmlFormatter_Characters(void* ptr, void* value);
void QXmlFormatter_Comment(void* ptr, char* value);
void QXmlFormatter_EndDocument(void* ptr);
void QXmlFormatter_EndElement(void* ptr);
void QXmlFormatter_EndOfSequence(void* ptr);
int QXmlFormatter_IndentationDepth(void* ptr);
void QXmlFormatter_ProcessingInstruction(void* ptr, void* name, char* value);
void QXmlFormatter_SetIndentationDepth(void* ptr, int depth);
void QXmlFormatter_StartDocument(void* ptr);
void QXmlFormatter_StartElement(void* ptr, void* name);
void QXmlFormatter_StartOfSequence(void* ptr);
void* QXmlItem_NewQXmlItem();
void* QXmlItem_NewQXmlItem4(void* atomicValue);
void* QXmlItem_NewQXmlItem2(void* other);
void* QXmlItem_NewQXmlItem3(void* node);
int QXmlItem_IsNode(void* ptr);
int QXmlItem_IsNull(void* ptr);
void QXmlItem_DestroyQXmlItem(void* ptr);
void* QXmlName_NewQXmlName();
void* QXmlName_NewQXmlName2(void* namePool, char* localName, char* namespaceURI, char* prefix);
int QXmlName_QXmlName_IsNCName(char* candidate);
int QXmlName_IsNull(void* ptr);
char* QXmlName_LocalName(void* ptr, void* namePool);
char* QXmlName_NamespaceUri(void* ptr, void* namePool);
char* QXmlName_Prefix(void* ptr, void* namePool);
char* QXmlName_ToClarkName(void* ptr, void* namePool);
void* QXmlNamePool_NewQXmlNamePool();
void* QXmlNamePool_NewQXmlNamePool2(void* other);
void QXmlNamePool_DestroyQXmlNamePool(void* ptr);
void* QXmlNodeModelIndex_NewQXmlNodeModelIndex();
void* QXmlNodeModelIndex_NewQXmlNodeModelIndex2(void* other);
void* QXmlNodeModelIndex_InternalPointer(void* ptr);
int QXmlNodeModelIndex_IsNull(void* ptr);
void* QXmlNodeModelIndex_Model(void* ptr);
void* QXmlQuery_NewQXmlQuery();
void* QXmlQuery_NewQXmlQuery4(int queryLanguage, void* np);
void* QXmlQuery_NewQXmlQuery3(void* np);
void* QXmlQuery_NewQXmlQuery2(void* other);
void QXmlQuery_BindVariable5(void* ptr, char* localName, void* device);
void QXmlQuery_BindVariable4(void* ptr, char* localName, void* value);
void QXmlQuery_BindVariable6(void* ptr, char* localName, void* query);
void QXmlQuery_BindVariable2(void* ptr, void* name, void* device);
void QXmlQuery_BindVariable(void* ptr, void* name, void* value);
void QXmlQuery_BindVariable3(void* ptr, void* name, void* query);
int QXmlQuery_IsValid(void* ptr);
void* QXmlQuery_MessageHandler(void* ptr);
void* QXmlQuery_NetworkAccessManager(void* ptr);
int QXmlQuery_QueryLanguage(void* ptr);
int QXmlQuery_SetFocus3(void* ptr, void* document);
int QXmlQuery_SetFocus4(void* ptr, char* focus);
int QXmlQuery_SetFocus2(void* ptr, void* documentURI);
void QXmlQuery_SetFocus(void* ptr, void* item);
void QXmlQuery_SetInitialTemplateName2(void* ptr, char* localName);
void QXmlQuery_SetInitialTemplateName(void* ptr, void* name);
void QXmlQuery_SetMessageHandler(void* ptr, void* aMessageHandler);
void QXmlQuery_SetNetworkAccessManager(void* ptr, void* newManager);
void QXmlQuery_SetQuery(void* ptr, void* sourceCode, void* documentURI);
void QXmlQuery_SetQuery3(void* ptr, char* sourceCode, void* documentURI);
void QXmlQuery_SetQuery2(void* ptr, void* queryURI, void* baseURI);
void QXmlQuery_SetUriResolver(void* ptr, void* resolver);
void* QXmlQuery_UriResolver(void* ptr);
void QXmlQuery_DestroyQXmlQuery(void* ptr);
void* QXmlResultItems_NewQXmlResultItems();
int QXmlResultItems_HasError(void* ptr);
void QXmlResultItems_DestroyQXmlResultItems(void* ptr);
void* QXmlSchema_NewQXmlSchema();
void* QXmlSchema_NewQXmlSchema2(void* other);
int QXmlSchema_IsValid(void* ptr);
int QXmlSchema_Load2(void* ptr, void* source, void* documentUri);
int QXmlSchema_Load3(void* ptr, void* data, void* documentUri);
int QXmlSchema_Load(void* ptr, void* source);
void* QXmlSchema_MessageHandler(void* ptr);
void* QXmlSchema_NetworkAccessManager(void* ptr);
void QXmlSchema_SetMessageHandler(void* ptr, void* handler);
void QXmlSchema_SetNetworkAccessManager(void* ptr, void* manager);
void QXmlSchema_SetUriResolver(void* ptr, void* resolver);
void* QXmlSchema_UriResolver(void* ptr);
void QXmlSchema_DestroyQXmlSchema(void* ptr);
void* QXmlSchemaValidator_NewQXmlSchemaValidator();
void* QXmlSchemaValidator_NewQXmlSchemaValidator2(void* schema);
void* QXmlSchemaValidator_MessageHandler(void* ptr);
void* QXmlSchemaValidator_NetworkAccessManager(void* ptr);
void QXmlSchemaValidator_SetMessageHandler(void* ptr, void* handler);
void QXmlSchemaValidator_SetNetworkAccessManager(void* ptr, void* manager);
void QXmlSchemaValidator_SetSchema(void* ptr, void* schema);
void QXmlSchemaValidator_SetUriResolver(void* ptr, void* resolver);
void* QXmlSchemaValidator_UriResolver(void* ptr);
int QXmlSchemaValidator_Validate2(void* ptr, void* source, void* documentUri);
int QXmlSchemaValidator_Validate3(void* ptr, void* data, void* documentUri);
int QXmlSchemaValidator_Validate(void* ptr, void* source);
void QXmlSchemaValidator_DestroyQXmlSchemaValidator(void* ptr);
void* QXmlSerializer_NewQXmlSerializer(void* query, void* outputDevice);
void QXmlSerializer_Attribute(void* ptr, void* name, void* value);
void QXmlSerializer_Characters(void* ptr, void* value);
void QXmlSerializer_Comment(void* ptr, char* value);
void QXmlSerializer_EndDocument(void* ptr);
void QXmlSerializer_EndElement(void* ptr);
void* QXmlSerializer_Codec(void* ptr);
void QXmlSerializer_EndOfSequence(void* ptr);
void QXmlSerializer_NamespaceBinding(void* ptr, void* nb);
void* QXmlSerializer_OutputDevice(void* ptr);
void QXmlSerializer_ProcessingInstruction(void* ptr, void* name, char* value);
void QXmlSerializer_SetCodec(void* ptr, void* outputCodec);
void QXmlSerializer_StartDocument(void* ptr);
void QXmlSerializer_StartElement(void* ptr, void* name);
void QXmlSerializer_StartOfSequence(void* ptr);

#ifdef __cplusplus
}
#endif