import 'dart:convert';
import 'dart:ffi';
import 'package:ffi/ffi.dart';

class GoSlice extends Struct {
  Pointer<Int64> data;

  @Int64()
  int len;

  @Int64()
  int cap;

  List<int> toList() {
    List<int> units = [];
    for (int i = 0; i < len; ++i) {
      units.add(data.elementAt(i).value);
    }
    return units;
  }

  static Pointer<GoSlice> fromList(List<int> units) {
    final ptr = allocate<Int64>(count: units.length);
    for (int i = 0; i < units.length; ++i) {
      ptr.elementAt(i).value = units[i];
    }
    final GoSlice slice = GoSlice();
    slice.data = ptr;
    slice.len = units.length;
    slice.cap = units.length;
    return slice.addressOf;
  }
}

class GoString extends Struct {
  Pointer<Uint8> string;

  @IntPtr()
  int length;

  String toString() {
    List<int> units = [];
    for (int i = 0; i < length; ++i) {
      units.add(string.elementAt(i).value);
    }
    return Utf8Decoder().convert(units);
  }

  static Pointer<GoString> fromString(String string) {
    List<int> units = Utf8Encoder().convert(string);
    final ptr = allocate<Uint8>(count: units.length);
    for (int i = 0; i < units.length; ++i) {
      ptr.elementAt(i).value = units[i];
    }
    final GoString str = allocate<GoString>().ref;
    str.length = units.length;
    str.string = ptr;
    return str.addressOf;
  }
}

typedef generate_keypair = Void Function(Pointer<GoString>);
typedef GenerateKeyPair = void Function(Pointer<GoString>);

void main(List<String> args) {
  final libsignal = DynamicLibrary.open("./libsignal.so");
  final generateKeyPair = libsignal.lookup<NativeFunction<generate_keypair>>("GenerateKeyPair").asFunction<GenerateKeyPair>();

  final Pointer<GoString> message = GoString.fromString("");
  generateKeyPair(message);
  print(message.ref);
  free(message);
}
