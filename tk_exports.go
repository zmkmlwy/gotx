package main

import (
	"reflect"

	"github.com/topxeq/tk"
)

func init() {
	GotxSymbols["github.com/topxeq/tk"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AESDecrypt":                               reflect.ValueOf(tk.AESDecrypt),
		"AESEncrypt":                               reflect.ValueOf(tk.AESEncrypt),
		"AddDebug":                                 reflect.ValueOf(tk.AddDebug),
		"AddDebugF":                                reflect.ValueOf(tk.AddDebugF),
		"AddLastSubString":                         reflect.ValueOf(tk.AddLastSubString),
		"AnalyzeCommandLineParamter":               reflect.ValueOf(tk.AnalyzeCommandLineParamter),
		"AnalyzeURLParams":                         reflect.ValueOf(tk.AnalyzeURLParams),
		"AppendDualLineList":                       reflect.ValueOf(tk.AppendDualLineList),
		"AppendSimpleMapFromFile":                  reflect.ValueOf(tk.AppendSimpleMapFromFile),
		"AppendStringToFile":                       reflect.ValueOf(tk.AppendStringToFile),
		"BitXor":                                   reflect.ValueOf(tk.BitXor),
		"ByteSliceToStringDec":                     reflect.ValueOf(tk.ByteSliceToStringDec),
		"ByteToHex":                                reflect.ValueOf(tk.ByteToHex),
		"BytesToHex":                               reflect.ValueOf(tk.BytesToHex),
		"CalCosineSimilarityBetweenFloatsBig":      reflect.ValueOf(tk.CalCosineSimilarityBetweenFloatsBig),
		"CheckErr":                                 reflect.ValueOf(tk.CheckErr),
		"CheckErrCompact":                          reflect.ValueOf(tk.CheckErrCompact),
		"CheckErrf":                                reflect.ValueOf(tk.CheckErrf),
		"CheckError":                               reflect.ValueOf(tk.CheckError),
		"CheckErrorFunc":                           reflect.ValueOf(tk.CheckErrorFunc),
		"CheckErrorString":                         reflect.ValueOf(tk.CheckErrorString),
		"ClearDebug":                               reflect.ValueOf(tk.ClearDebug),
		"Contains":                                 reflect.ValueOf(tk.Contains),
		"ContainsIgnoreCase":                       reflect.ValueOf(tk.ContainsIgnoreCase),
		"ContainsIn":                               reflect.ValueOf(tk.ContainsIn),
		"ContainsInStringList":                     reflect.ValueOf(tk.ContainsInStringList),
		"ConvertStringToUTF8":                      reflect.ValueOf(tk.ConvertStringToUTF8),
		"ConvertToGB18030":                         reflect.ValueOf(tk.ConvertToGB18030),
		"ConvertToGB18030Bytes":                    reflect.ValueOf(tk.ConvertToGB18030Bytes),
		"ConvertToUTF8":                            reflect.ValueOf(tk.ConvertToUTF8),
		"CopyFile":                                 reflect.ValueOf(tk.CopyFile),
		"CreateSimpleEvent":                        reflect.ValueOf(tk.CreateSimpleEvent),
		"CreateString":                             reflect.ValueOf(tk.CreateString),
		"CreateStringEmpty":                        reflect.ValueOf(tk.CreateStringEmpty),
		"CreateStringError":                        reflect.ValueOf(tk.CreateStringError),
		"CreateStringErrorF":                       reflect.ValueOf(tk.CreateStringErrorF),
		"CreateStringErrorFromTXError":             reflect.ValueOf(tk.CreateStringErrorFromTXError),
		"CreateStringSimple":                       reflect.ValueOf(tk.CreateStringSimple),
		"CreateStringSuccess":                      reflect.ValueOf(tk.CreateStringSuccess),
		"CreateStringWithObject":                   reflect.ValueOf(tk.CreateStringWithObject),
		"CreateTXCollection":                       reflect.ValueOf(tk.CreateTXCollection),
		"CreateTempFile":                           reflect.ValueOf(tk.CreateTempFile),
		"DebugModeG":                               reflect.ValueOf(&tk.DebugModeG).Elem(),
		"DecodeStringCustom":                       reflect.ValueOf(tk.DecodeStringCustom),
		"DecodeStringSimple":                       reflect.ValueOf(tk.DecodeStringSimple),
		"DecodeStringUnderline":                    reflect.ValueOf(tk.DecodeStringUnderline),
		"DecryptDataByTXDEE":                       reflect.ValueOf(tk.DecryptDataByTXDEE),
		"DecryptDataByTXDEF":                       reflect.ValueOf(tk.DecryptDataByTXDEF),
		"DecryptFileByTXDEF":                       reflect.ValueOf(tk.DecryptFileByTXDEF),
		"DecryptFileByTXDEFS":                      reflect.ValueOf(tk.DecryptFileByTXDEFS),
		"DecryptFileByTXDEFStream":                 reflect.ValueOf(tk.DecryptFileByTXDEFStream),
		"DecryptFileByTXDEFStreamS":                reflect.ValueOf(tk.DecryptFileByTXDEFStreamS),
		"DecryptStreamByTXDEF":                     reflect.ValueOf(tk.DecryptStreamByTXDEF),
		"DecryptStringByTXDEE":                     reflect.ValueOf(tk.DecryptStringByTXDEE),
		"DecryptStringByTXDEF":                     reflect.ValueOf(tk.DecryptStringByTXDEF),
		"DecryptStringByTXTE":                      reflect.ValueOf(tk.DecryptStringByTXTE),
		"DeepClone":                                reflect.ValueOf(tk.DeepClone),
		"DeepCopyFromTo":                           reflect.ValueOf(tk.DeepCopyFromTo),
		"DeleteItemInInt64Array":                   reflect.ValueOf(tk.DeleteItemInInt64Array),
		"DeleteItemInIntArray":                     reflect.ValueOf(tk.DeleteItemInIntArray),
		"DeleteItemInStringArray":                  reflect.ValueOf(tk.DeleteItemInStringArray),
		"DownloadBytes":                            reflect.ValueOf(tk.DownloadBytes),
		"DownloadFile":                             reflect.ValueOf(tk.DownloadFile),
		"DownloadPage":                             reflect.ValueOf(tk.DownloadPage),
		"DownloadPageByMap":                        reflect.ValueOf(tk.DownloadPageByMap),
		"DownloadPageUTF8":                         reflect.ValueOf(tk.DownloadPageUTF8),
		"EncodeStringCustom":                       reflect.ValueOf(tk.EncodeStringCustom),
		"EncodeStringCustomEx":                     reflect.ValueOf(tk.EncodeStringCustomEx),
		"EncodeStringSimple":                       reflect.ValueOf(tk.EncodeStringSimple),
		"EncodeStringUnderline":                    reflect.ValueOf(tk.EncodeStringUnderline),
		"EncodeToXMLString":                        reflect.ValueOf(tk.EncodeToXMLString),
		"EncryptDataByTXDEE":                       reflect.ValueOf(tk.EncryptDataByTXDEE),
		"EncryptDataByTXDEF":                       reflect.ValueOf(tk.EncryptDataByTXDEF),
		"EncryptFileByTXDEF":                       reflect.ValueOf(tk.EncryptFileByTXDEF),
		"EncryptFileByTXDEFS":                      reflect.ValueOf(tk.EncryptFileByTXDEFS),
		"EncryptFileByTXDEFStream":                 reflect.ValueOf(tk.EncryptFileByTXDEFStream),
		"EncryptFileByTXDEFStreamS":                reflect.ValueOf(tk.EncryptFileByTXDEFStreamS),
		"EncryptFileByTXDEFWithHeader":             reflect.ValueOf(tk.EncryptFileByTXDEFWithHeader),
		"EncryptStreamByTXDEF":                     reflect.ValueOf(tk.EncryptStreamByTXDEF),
		"EncryptStringByTXDEE":                     reflect.ValueOf(tk.EncryptStringByTXDEE),
		"EncryptStringByTXDEF":                     reflect.ValueOf(tk.EncryptStringByTXDEF),
		"EncryptStringByTXTE":                      reflect.ValueOf(tk.EncryptStringByTXTE),
		"EndsWith":                                 reflect.ValueOf(tk.EndsWith),
		"EndsWithIgnoreCase":                       reflect.ValueOf(tk.EndsWithIgnoreCase),
		"EnsureBasePath":                           reflect.ValueOf(tk.EnsureBasePath),
		"EnsureMakeDirs":                           reflect.ValueOf(tk.EnsureMakeDirs),
		"EnsureMakeDirsE":                          reflect.ValueOf(tk.EnsureMakeDirsE),
		"EnsureValidFileNameX":                     reflect.ValueOf(tk.EnsureValidFileNameX),
		"ErrStrF":                                  reflect.ValueOf(tk.ErrStrF),
		"ErrToStr":                                 reflect.ValueOf(tk.ErrToStr),
		"ErrToStrF":                                reflect.ValueOf(tk.ErrToStrF),
		"Errf":                                     reflect.ValueOf(tk.Errf),
		"ErrorStringToError":                       reflect.ValueOf(tk.ErrorStringToError),
		"ErrorToString":                            reflect.ValueOf(tk.ErrorToString),
		"Exit":                                     reflect.ValueOf(tk.Exit),
		"FatalErr":                                 reflect.ValueOf(tk.FatalErr),
		"FatalErrf":                                reflect.ValueOf(tk.FatalErrf),
		"Fatalf":                                   reflect.ValueOf(tk.Fatalf),
		"FindFirstDiffIndex":                       reflect.ValueOf(tk.FindFirstDiffIndex),
		"FindSamePrefix":                           reflect.ValueOf(tk.FindSamePrefix),
		"FlattenXML":                               reflect.ValueOf(tk.FlattenXML),
		"Float32ArrayToFloat64Array":               reflect.ValueOf(tk.Float32ArrayToFloat64Array),
		"Float64ToStr":                             reflect.ValueOf(tk.Float64ToStr),
		"FormatStringSliceSlice":                   reflect.ValueOf(tk.FormatStringSliceSlice),
		"FormatTime":                               reflect.ValueOf(tk.FormatTime),
		"Fpl":                                      reflect.ValueOf(tk.Fpl),
		"Fpr":                                      reflect.ValueOf(tk.Fpr),
		"FromJSON":                                 reflect.ValueOf(tk.FromJSON),
		"FromJSONWithDefault":                      reflect.ValueOf(tk.FromJSONWithDefault),
		"GenerateErrorString":                      reflect.ValueOf(tk.GenerateErrorString),
		"GenerateErrorStringF":                     reflect.ValueOf(tk.GenerateErrorStringF),
		"GenerateErrorStringFTX":                   reflect.ValueOf(tk.GenerateErrorStringFTX),
		"GenerateErrorStringTX":                    reflect.ValueOf(tk.GenerateErrorStringTX),
		"GenerateFileListInDir":                    reflect.ValueOf(tk.GenerateFileListInDir),
		"GenerateFileListRecursively":              reflect.ValueOf(tk.GenerateFileListRecursively),
		"GenerateFileListRecursivelyWithExclusive": reflect.ValueOf(tk.GenerateFileListRecursivelyWithExclusive),
		"GenerateJSONPResponse":                    reflect.ValueOf(tk.GenerateJSONPResponse),
		"GenerateJSONPResponseWith2Object":         reflect.ValueOf(tk.GenerateJSONPResponseWith2Object),
		"GenerateJSONPResponseWith3Object":         reflect.ValueOf(tk.GenerateJSONPResponseWith3Object),
		"GenerateJSONPResponseWithObject":          reflect.ValueOf(tk.GenerateJSONPResponseWithObject),
		"GenerateRandomString":                     reflect.ValueOf(tk.GenerateRandomString),
		"GetAllOSParameters":                       reflect.ValueOf(tk.GetAllOSParameters),
		"GetAllParameters":                         reflect.ValueOf(tk.GetAllParameters),
		"GetAllSwitches":                           reflect.ValueOf(tk.GetAllSwitches),
		"GetApplicationPath":                       reflect.ValueOf(tk.GetApplicationPath),
		"GetAvailableFileName":                     reflect.ValueOf(tk.GetAvailableFileName),
		"GetClipText":                              reflect.ValueOf(tk.GetClipText),
		"GetCurrentDir":                            reflect.ValueOf(tk.GetCurrentDir),
		"GetCurrentThreadID":                       reflect.ValueOf(tk.GetCurrentThreadID),
		"GetDBConnection":                          reflect.ValueOf(tk.GetDBConnection),
		"GetDBResultArray":                         reflect.ValueOf(tk.GetDBResultArray),
		"GetDBResultString":                        reflect.ValueOf(tk.GetDBResultString),
		"GetDBRowCount":                            reflect.ValueOf(tk.GetDBRowCount),
		"GetDBRowCountCompact":                     reflect.ValueOf(tk.GetDBRowCountCompact),
		"GetDebug":                                 reflect.ValueOf(tk.GetDebug),
		"GetDirOfFilePath":                         reflect.ValueOf(tk.GetDirOfFilePath),
		"GetEnv":                                   reflect.ValueOf(tk.GetEnv),
		"GetErrorString":                           reflect.ValueOf(tk.GetErrorString),
		"GetErrorStringSafely":                     reflect.ValueOf(tk.GetErrorStringSafely),
		"GetFileExt":                               reflect.ValueOf(tk.GetFileExt),
		"GetFilePathSeperator":                     reflect.ValueOf(tk.GetFilePathSeperator),
		"GetFormValueWithDefaultValue":             reflect.ValueOf(tk.GetFormValueWithDefaultValue),
		"GetGlobalEnvList":                         reflect.ValueOf(tk.GetGlobalEnvList),
		"GetGlobalEnvString":                       reflect.ValueOf(tk.GetGlobalEnvString),
		"GetInputBufferedScan":                     reflect.ValueOf(tk.GetInputBufferedScan),
		"GetInputPasswordf":                        reflect.ValueOf(tk.GetInputPasswordf),
		"GetInputf":                                reflect.ValueOf(tk.GetInputf),
		"GetJSONNode":                              reflect.ValueOf(tk.GetJSONNode),
		"GetJSONNodeAny":                           reflect.ValueOf(tk.GetJSONNodeAny),
		"GetJSONSubNode":                           reflect.ValueOf(tk.GetJSONSubNode),
		"GetJSONSubNodeAny":                        reflect.ValueOf(tk.GetJSONSubNodeAny),
		"GetLastComponentOfFilePath":               reflect.ValueOf(tk.GetLastComponentOfFilePath),
		"GetLastComponentOfUrl":                    reflect.ValueOf(tk.GetLastComponentOfUrl),
		"GetLoginAuth":                             reflect.ValueOf(tk.GetLoginAuth),
		"GetMSIStringWithDefault":                  reflect.ValueOf(tk.GetMSIStringWithDefault),
		"GetMSSArrayFromXML":                       reflect.ValueOf(tk.GetMSSArrayFromXML),
		"GetMSSFromXML":                            reflect.ValueOf(tk.GetMSSFromXML),
		"GetNowDateString":                         reflect.ValueOf(tk.GetNowDateString),
		"GetNowMinutesInDay":                       reflect.ValueOf(tk.GetNowMinutesInDay),
		"GetNowTimeOnlyStringBeijing":              reflect.ValueOf(tk.GetNowTimeOnlyStringBeijing),
		"GetNowTimeString":                         reflect.ValueOf(tk.GetNowTimeString),
		"GetNowTimeStringFormal":                   reflect.ValueOf(tk.GetNowTimeStringFormal),
		"GetNowTimeStringFormat":                   reflect.ValueOf(tk.GetNowTimeStringFormat),
		"GetNowTimeStringHourMinute":               reflect.ValueOf(tk.GetNowTimeStringHourMinute),
		"GetOSArgs":                                reflect.ValueOf(tk.GetOSArgs),
		"GetOSName":                                reflect.ValueOf(tk.GetOSName),
		"GetParameterByIndexWithDefaultValue":      reflect.ValueOf(tk.GetParameterByIndexWithDefaultValue),
		"GetPlainAuth":                             reflect.ValueOf(tk.GetPlainAuth),
		"GetRandomByte":                            reflect.ValueOf(tk.GetRandomByte),
		"GetRandomInt64InRange":                    reflect.ValueOf(tk.GetRandomInt64InRange),
		"GetRandomInt64LessThan":                   reflect.ValueOf(tk.GetRandomInt64LessThan),
		"GetRandomIntInRange":                      reflect.ValueOf(tk.GetRandomIntInRange),
		"GetRandomIntLessThan":                     reflect.ValueOf(tk.GetRandomIntLessThan),
		"GetRandomSubDualList":                     reflect.ValueOf(tk.GetRandomSubDualList),
		"GetRandomizeInt64ArrayCopy":               reflect.ValueOf(tk.GetRandomizeInt64ArrayCopy),
		"GetRandomizeIntArrayCopy":                 reflect.ValueOf(tk.GetRandomizeIntArrayCopy),
		"GetRandomizeStringArrayCopy":              reflect.ValueOf(tk.GetRandomizeStringArrayCopy),
		"GetRuntimeStack":                          reflect.ValueOf(tk.GetRuntimeStack),
		"GetSliceMaxLen":                           reflect.ValueOf(tk.GetSliceMaxLen),
		"GetStringSliceFilled":                     reflect.ValueOf(tk.GetStringSliceFilled),
		"GetSuccessValue":                          reflect.ValueOf(tk.GetSuccessValue),
		"GetSwitchWithDefaultInt64Value":           reflect.ValueOf(tk.GetSwitchWithDefaultInt64Value),
		"GetSwitchWithDefaultIntValue":             reflect.ValueOf(tk.GetSwitchWithDefaultIntValue),
		"GetSwitchWithDefaultValue":                reflect.ValueOf(tk.GetSwitchWithDefaultValue),
		"GetTimeFromUnixTimeStamp":                 reflect.ValueOf(tk.GetTimeFromUnixTimeStamp),
		"GetTimeFromUnixTimeStampMid":              reflect.ValueOf(tk.GetTimeFromUnixTimeStampMid),
		"GetTimeStamp":                             reflect.ValueOf(tk.GetTimeStamp),
		"GetTimeStampMid":                          reflect.ValueOf(tk.GetTimeStampMid),
		"GetTimeStampNano":                         reflect.ValueOf(tk.GetTimeStampNano),
		"GetTimeStringDiffMS":                      reflect.ValueOf(tk.GetTimeStringDiffMS),
		"GetUserInput":                             reflect.ValueOf(tk.GetUserInput),
		"GetValue":                                 reflect.ValueOf(tk.GetValue),
		"GetValueOfMSS":                            reflect.ValueOf(tk.GetValueOfMSS),
		"GetVar":                                   reflect.ValueOf(tk.GetVar),
		"GlobalEnvSetG":                            reflect.ValueOf(&tk.GlobalEnvSetG).Elem(),
		"HasGlobalEnv":                             reflect.ValueOf(tk.HasGlobalEnv),
		"HexToBytes":                               reflect.ValueOf(tk.HexToBytes),
		"HexToInt":                                 reflect.ValueOf(tk.HexToInt),
		"IfFileExists":                             reflect.ValueOf(tk.IfFileExists),
		"IfSwitchExists":                           reflect.ValueOf(tk.IfSwitchExists),
		"IfSwitchExistsWhole":                      reflect.ValueOf(tk.IfSwitchExistsWhole),
		"InStrings":                                reflect.ValueOf(tk.InStrings),
		"IndexInStringList":                        reflect.ValueOf(tk.IndexInStringList),
		"IndexInStringListFromEnd":                 reflect.ValueOf(tk.IndexInStringListFromEnd),
		"Int64ArrayToFloat64Array":                 reflect.ValueOf(tk.Int64ArrayToFloat64Array),
		"Int64ToStr":                               reflect.ValueOf(tk.Int64ToStr),
		"IntToKMGT":                                reflect.ValueOf(tk.IntToKMGT),
		"IntToStr":                                 reflect.ValueOf(tk.IntToStr),
		"IntToWYZ":                                 reflect.ValueOf(tk.IntToWYZ),
		"IsDirectory":                              reflect.ValueOf(tk.IsDirectory),
		"IsEmptyTrim":                              reflect.ValueOf(tk.IsEmptyTrim),
		"IsErrorString":                            reflect.ValueOf(tk.IsErrorString),
		"IsFile":                                   reflect.ValueOf(tk.IsFile),
		"IsFloat64NearlyEqual":                     reflect.ValueOf(tk.IsFloat64NearlyEqual),
		"IsValidEmail":                             reflect.ValueOf(tk.IsValidEmail),
		"IsYesterday":                              reflect.ValueOf(tk.IsYesterday),
		"JSONToMapStringString":                    reflect.ValueOf(tk.JSONToMapStringString),
		"JSONToObject":                             reflect.ValueOf(tk.JSONToObject),
		"JSONToObjectE":                            reflect.ValueOf(tk.JSONToObjectE),
		"JSONToStringArray":                        reflect.ValueOf(tk.JSONToStringArray),
		"JoinDualList":                             reflect.ValueOf(tk.JoinDualList),
		"JoinLines":                                reflect.ValueOf(tk.JoinLines),
		"JoinLinesBySeparator":                     reflect.ValueOf(tk.JoinLinesBySeparator),
		"JoinPath":                                 reflect.ValueOf(tk.JoinPath),
		"JoinURL":                                  reflect.ValueOf(tk.JoinURL),
		"KindOfValueReflect":                       reflect.ValueOf(tk.KindOfValueReflect),
		"Len64":                                    reflect.ValueOf(tk.Len64),
		"LoadBytes":                                reflect.ValueOf(tk.LoadBytes),
		"LoadBytesFromFileE":                       reflect.ValueOf(tk.LoadBytesFromFileE),
		"LoadDualLineList":                         reflect.ValueOf(tk.LoadDualLineList),
		"LoadDualLineListFromString":               reflect.ValueOf(tk.LoadDualLineListFromString),
		"LoadSimpleMapFromDir":                     reflect.ValueOf(tk.LoadSimpleMapFromDir),
		"LoadSimpleMapFromFile":                    reflect.ValueOf(tk.LoadSimpleMapFromFile),
		"LoadSimpleMapFromFileE":                   reflect.ValueOf(tk.LoadSimpleMapFromFileE),
		"LoadSimpleMapFromString":                  reflect.ValueOf(tk.LoadSimpleMapFromString),
		"LoadSimpleMapFromStringE":                 reflect.ValueOf(tk.LoadSimpleMapFromStringE),
		"LoadStringFromFile":                       reflect.ValueOf(tk.LoadStringFromFile),
		"LoadStringFromFileB":                      reflect.ValueOf(tk.LoadStringFromFileB),
		"LoadStringFromFileE":                      reflect.ValueOf(tk.LoadStringFromFileE),
		"LoadStringFromFileWithDefault":            reflect.ValueOf(tk.LoadStringFromFileWithDefault),
		"LoadStringList":                           reflect.ValueOf(tk.LoadStringList),
		"LoadStringListBuffered":                   reflect.ValueOf(tk.LoadStringListBuffered),
		"LoadStringListFromFile":                   reflect.ValueOf(tk.LoadStringListFromFile),
		"LoadStringTX":                             reflect.ValueOf(tk.LoadStringTX),
		"LogWithTime":                              reflect.ValueOf(tk.LogWithTime),
		"LogWithTimeCompact":                       reflect.ValueOf(tk.LogWithTimeCompact),
		"MD5Encrypt":                               reflect.ValueOf(tk.MD5Encrypt),
		"NewRandomGenerator":                       reflect.ValueOf(tk.NewRandomGenerator),
		"NewSSHClient":                             reflect.ValueOf(tk.NewSSHClient),
		"NowToFileName":                            reflect.ValueOf(tk.NowToFileName),
		"NowToStrUTC":                              reflect.ValueOf(tk.NowToStrUTC),
		"ObjectToJSON":                             reflect.ValueOf(tk.ObjectToJSON),
		"ObjectToJSONIndent":                       reflect.ValueOf(tk.ObjectToJSONIndent),
		"ParseCommandLine":                         reflect.ValueOf(tk.ParseCommandLine),
		"ParseHexColor":                            reflect.ValueOf(tk.ParseHexColor),
		"Pkcs7Padding":                             reflect.ValueOf(tk.Pkcs7Padding),
		"Pl":                                       reflect.ValueOf(tk.Pl),
		"PlAndExit":                                reflect.ValueOf(tk.PlAndExit),
		"PlErr":                                    reflect.ValueOf(tk.PlErr),
		"PlErrAndExit":                             reflect.ValueOf(tk.PlErrAndExit),
		"PlErrSimple":                              reflect.ValueOf(tk.PlErrSimple),
		"PlErrSimpleAndExit":                       reflect.ValueOf(tk.PlErrSimpleAndExit),
		"PlErrString":                              reflect.ValueOf(tk.PlErrString),
		"PlErrWithPrefix":                          reflect.ValueOf(tk.PlErrWithPrefix),
		"PlNow":                                    reflect.ValueOf(tk.PlNow),
		"PlSimpleErrorString":                      reflect.ValueOf(tk.PlSimpleErrorString),
		"PlTXErr":                                  reflect.ValueOf(tk.PlTXErr),
		"PlVerbose":                                reflect.ValueOf(tk.PlVerbose),
		"Pln":                                      reflect.ValueOf(tk.Pln),
		"Plv":                                      reflect.ValueOf(tk.Plv),
		"PlvWithError":                             reflect.ValueOf(tk.PlvWithError),
		"Plvs":                                     reflect.ValueOf(tk.Plvs),
		"Plvsr":                                    reflect.ValueOf(tk.Plvsr),
		"PostRequest":                              reflect.ValueOf(tk.PostRequest),
		"PostRequestBytesWithCookieX":              reflect.ValueOf(tk.PostRequestBytesWithCookieX),
		"PostRequestBytesWithMSSHeaderX":           reflect.ValueOf(tk.PostRequestBytesWithMSSHeaderX),
		"PostRequestBytesX":                        reflect.ValueOf(tk.PostRequestBytesX),
		"PostRequestX":                             reflect.ValueOf(tk.PostRequestX),
		"Pr":                                       reflect.ValueOf(tk.Pr),
		"Prf":                                      reflect.ValueOf(tk.Prf),
		"Printf":                                   reflect.ValueOf(tk.Printf),
		"Printfln":                                 reflect.ValueOf(tk.Printfln),
		"Prl":                                      reflect.ValueOf(tk.Prl),
		"Randomize":                                reflect.ValueOf(tk.Randomize),
		"ReadLineFromBufioReader":                  reflect.ValueOf(tk.ReadLineFromBufioReader),
		"RegContains":                              reflect.ValueOf(tk.RegContains),
		"RegFindAll":                               reflect.ValueOf(tk.RegFindAll),
		"RegFindAllGroups":                         reflect.ValueOf(tk.RegFindAllGroups),
		"RegFindFirst":                             reflect.ValueOf(tk.RegFindFirst),
		"RegFindFirstIndex":                        reflect.ValueOf(tk.RegFindFirstIndex),
		"RegFindFirstTX":                           reflect.ValueOf(tk.RegFindFirstTX),
		"RegMatch":                                 reflect.ValueOf(tk.RegMatch),
		"RegReplace":                               reflect.ValueOf(tk.RegReplace),
		"RegStartsWith":                            reflect.ValueOf(tk.RegStartsWith),
		"RemoveBOM":                                reflect.ValueOf(tk.RemoveBOM),
		"RemoveDuplicateInDualLineList":            reflect.ValueOf(tk.RemoveDuplicateInDualLineList),
		"RemoveFile":                               reflect.ValueOf(tk.RemoveFile),
		"RemoveFileExt":                            reflect.ValueOf(tk.RemoveFileExt),
		"RemoveGlobalEnv":                          reflect.ValueOf(tk.RemoveGlobalEnv),
		"RemoveItemsInArray":                       reflect.ValueOf(tk.RemoveItemsInArray),
		"RemoveLastSubString":                      reflect.ValueOf(tk.RemoveLastSubString),
		"Replace":                                  reflect.ValueOf(tk.Replace),
		"ReplaceLineEnds":                          reflect.ValueOf(tk.ReplaceLineEnds),
		"ReshapeXML":                               reflect.ValueOf(tk.ReshapeXML),
		"RestoreLineEnds":                          reflect.ValueOf(tk.RestoreLineEnds),
		"RunWinFileWithSystemDefault":              reflect.ValueOf(tk.RunWinFileWithSystemDefault),
		"SafelyGetFloat64ForKeyWithDefault":        reflect.ValueOf(tk.SafelyGetFloat64ForKeyWithDefault),
		"SafelyGetIntForKeyWithDefault":            reflect.ValueOf(tk.SafelyGetIntForKeyWithDefault),
		"SafelyGetStringForKeyWithDefault":         reflect.ValueOf(tk.SafelyGetStringForKeyWithDefault),
		"SaveDualLineList":                         reflect.ValueOf(tk.SaveDualLineList),
		"SaveSimpleMapToFile":                      reflect.ValueOf(tk.SaveSimpleMapToFile),
		"SaveStringList":                           reflect.ValueOf(tk.SaveStringList),
		"SaveStringListBuffered":                   reflect.ValueOf(tk.SaveStringListBuffered),
		"SaveStringListBufferedByRange":            reflect.ValueOf(tk.SaveStringListBufferedByRange),
		"SaveStringListWin":                        reflect.ValueOf(tk.SaveStringListWin),
		"SaveStringToFile":                         reflect.ValueOf(tk.SaveStringToFile),
		"SaveStringToFileE":                        reflect.ValueOf(tk.SaveStringToFileE),
		"SetClipText":                              reflect.ValueOf(tk.SetClipText),
		"SetGlobalEnv":                             reflect.ValueOf(tk.SetGlobalEnv),
		"SetLogFile":                               reflect.ValueOf(tk.SetLogFile),
		"SetValue":                                 reflect.ValueOf(tk.SetValue),
		"SetVar":                                   reflect.ValueOf(tk.SetVar),
		"ShuffleStringArray":                       reflect.ValueOf(tk.ShuffleStringArray),
		"SimpleMapToString":                        reflect.ValueOf(tk.SimpleMapToString),
		"SleepMilliSeconds":                        reflect.ValueOf(tk.SleepMilliSeconds),
		"SleepSeconds":                             reflect.ValueOf(tk.SleepSeconds),
		"Split":                                    reflect.ValueOf(tk.Split),
		"SplitLines":                               reflect.ValueOf(tk.SplitLines),
		"SplitLinesRemoveEmpty":                    reflect.ValueOf(tk.SplitLinesRemoveEmpty),
		"SplitN":                                   reflect.ValueOf(tk.SplitN),
		"Spr":                                      reflect.ValueOf(tk.Spr),
		"StartsWith":                               reflect.ValueOf(tk.StartsWith),
		"StartsWithBOM":                            reflect.ValueOf(tk.StartsWithBOM),
		"StartsWithDigit":                          reflect.ValueOf(tk.StartsWithDigit),
		"StartsWithIgnoreCase":                     reflect.ValueOf(tk.StartsWithIgnoreCase),
		"StartsWithUpper":                          reflect.ValueOf(tk.StartsWithUpper),
		"StrToBool":                                reflect.ValueOf(tk.StrToBool),
		"StrToFloat64":                             reflect.ValueOf(tk.StrToFloat64),
		"StrToFloat64WithDefaultValue":             reflect.ValueOf(tk.StrToFloat64WithDefaultValue),
		"StrToInt":                                 reflect.ValueOf(tk.StrToInt),
		"StrToInt64WithDefaultValue":               reflect.ValueOf(tk.StrToInt64WithDefaultValue),
		"StrToIntPositive":                         reflect.ValueOf(tk.StrToIntPositive),
		"StrToIntWithDefaultValue":                 reflect.ValueOf(tk.StrToIntWithDefaultValue),
		"StrToTime":                                reflect.ValueOf(tk.StrToTime),
		"StrToTimeByFormat":                        reflect.ValueOf(tk.StrToTimeByFormat),
		"StrToTimeCompact":                         reflect.ValueOf(tk.StrToTimeCompact),
		"StrToTimeCompactNoError":                  reflect.ValueOf(tk.StrToTimeCompactNoError),
		"StringReplace":                            reflect.ValueOf(tk.StringReplace),
		"SumBytes":                                 reflect.ValueOf(tk.SumBytes),
		"SystemCmd":                                reflect.ValueOf(tk.SystemCmd),
		"TXDEF_BUFFER_LEN":                         reflect.ValueOf(tk.TXDEF_BUFFER_LEN),
		"TXResultFromString":                       reflect.ValueOf(tk.TXResultFromString),
		"TXResultFromStringE":                      reflect.ValueOf(tk.TXResultFromStringE),
		"TimeFormat":                               reflect.ValueOf(&tk.TimeFormat).Elem(),
		"TimeFormatCompact":                        reflect.ValueOf(&tk.TimeFormatCompact).Elem(),
		"TimeFormatCompact2":                       reflect.ValueOf(&tk.TimeFormatCompact2).Elem(),
		"TimeFormatMS":                             reflect.ValueOf(&tk.TimeFormatMS).Elem(),
		"ToJSON":                                   reflect.ValueOf(tk.ToJSON),
		"ToJSONIndent":                             reflect.ValueOf(tk.ToJSONIndent),
		"ToJSONIndentWithDefault":                  reflect.ValueOf(tk.ToJSONIndentWithDefault),
		"ToJSONWithDefault":                        reflect.ValueOf(tk.ToJSONWithDefault),
		"ToLower":                                  reflect.ValueOf(tk.ToLower),
		"ToUpper":                                  reflect.ValueOf(tk.ToUpper),
		"Trim":                                     reflect.ValueOf(tk.Trim),
		"TrimCharSet":                              reflect.ValueOf(tk.TrimCharSet),
		"TypeOfValue":                              reflect.ValueOf(tk.TypeOfValue),
		"TypeOfValueReflect":                       reflect.ValueOf(tk.TypeOfValueReflect),
		"UrlDecode":                                reflect.ValueOf(tk.UrlDecode),
		"UrlEncode":                                reflect.ValueOf(tk.UrlEncode),
		"UrlEncode2":                               reflect.ValueOf(tk.UrlEncode2),
		"MSSFromJSON":                              reflect.ValueOf(tk.MSSFromJSON),

		// type definitions
		"ExitCallback":  reflect.ValueOf((*tk.ExitCallback)(nil)),
		"LoginAuth":     reflect.ValueOf((*tk.LoginAuth)(nil)),
		"PlainAuth":     reflect.ValueOf((*tk.PlainAuth)(nil)),
		"RandomX":       reflect.ValueOf((*tk.RandomX)(nil)),
		"ServerInfo":    reflect.ValueOf((*tk.ServerInfo)(nil)),
		"SimpleEvent":   reflect.ValueOf((*tk.SimpleEvent)(nil)),
		"TXCollection":  reflect.ValueOf((*tk.TXCollection)(nil)),
		"TXResult":      reflect.ValueOf((*tk.TXResult)(nil)),
		"TXString":      reflect.ValueOf((*tk.TXString)(nil)),
		"TXStringArray": reflect.ValueOf((*tk.TXStringArray)(nil)),
		"TXStringSlice": reflect.ValueOf((*tk.TXStringSlice)(nil)),
	}

	GotxSymbols["tk"] = GotxSymbols["github.com/topxeq/tk"]
}
