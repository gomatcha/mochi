// Generated by Cap'n Proto compiler, DO NOT EDIT
// source: view.capnp

#ifndef CAPNP_INCLUDED_9abec47167560884_
#define CAPNP_INCLUDED_9abec47167560884_

#include <capnp/generated-header-support.h>

#if CAPNP_VERSION != 5003
#error "Version mismatch between generated code and library headers.  You must use the same version of the Cap'n Proto compiler and library."
#endif

#include <layout/encoding/layout.capnp.h>

namespace capnp {
namespace schemas {

CAPNP_DECLARE_SCHEMA(e94009e96d6eeb30);
CAPNP_DECLARE_SCHEMA(83927b3747bee5da);
CAPNP_DECLARE_SCHEMA(bd0993cece143673);
CAPNP_DECLARE_SCHEMA(91412d97d3671d43);

}  // namespace schemas
}  // namespace capnp


struct Node {
  Node() = delete;

  class Reader;
  class Builder;
  class Pipeline;

  struct _capnpPrivate {
    CAPNP_DECLARE_STRUCT_HEADER(e94009e96d6eeb30, 4, 3)
    #if !CAPNP_LITE
    static constexpr ::capnp::_::RawBrandedSchema const* brand = &schema->defaultBrand;
    #endif  // !CAPNP_LITE
  };
};

struct Root {
  Root() = delete;

  class Reader;
  class Builder;
  class Pipeline;

  struct _capnpPrivate {
    CAPNP_DECLARE_STRUCT_HEADER(83927b3747bee5da, 0, 1)
    #if !CAPNP_LITE
    static constexpr ::capnp::_::RawBrandedSchema const* brand = &schema->defaultBrand;
    #endif  // !CAPNP_LITE
  };
};

template <typename Key = ::capnp::AnyPointer, typename Value = ::capnp::AnyPointer>
struct Map {
  Map() = delete;

  class Reader;
  class Builder;
  class Pipeline;
  struct Entry;

  struct _capnpPrivate {
    CAPNP_DECLARE_STRUCT_HEADER(bd0993cece143673, 0, 1)
    #if !CAPNP_LITE
    static const ::capnp::_::RawBrandedSchema::Scope brandScopes[];
    static const ::capnp::_::RawBrandedSchema::Binding brandBindings[];
    static const ::capnp::_::RawBrandedSchema::Dependency brandDependencies[];
    static const ::capnp::_::RawBrandedSchema specificBrand;
    static constexpr ::capnp::_::RawBrandedSchema const* brand = ::capnp::_::ChooseBrand<_capnpPrivate, Key, Value>::brand;
    #endif  // !CAPNP_LITE
  };
};

template <typename Key, typename Value>
struct Map<Key, Value>::Entry {
  Entry() = delete;

  class Reader;
  class Builder;
  class Pipeline;

  struct _capnpPrivate {
    CAPNP_DECLARE_STRUCT_HEADER(91412d97d3671d43, 0, 2)
    #if !CAPNP_LITE
    static const ::capnp::_::RawBrandedSchema::Scope brandScopes[];
    static const ::capnp::_::RawBrandedSchema::Binding brandBindings[];
    static const ::capnp::_::RawBrandedSchema specificBrand;
    static constexpr ::capnp::_::RawBrandedSchema const* brand = ::capnp::_::ChooseBrand<_capnpPrivate, Key, Value>::brand;
    #endif  // !CAPNP_LITE
  };
};

// =======================================================================================

class Node::Reader {
public:
  typedef Node Reads;

  Reader() = default;
  inline explicit Reader(::capnp::_::StructReader base): _reader(base) {}

  inline ::capnp::MessageSize totalSize() const {
    return _reader.totalSize().asPublic();
  }

#if !CAPNP_LITE
  inline ::kj::StringTree toString() const {
    return ::capnp::_::structString(_reader, *_capnpPrivate::brand);
  }
#endif  // !CAPNP_LITE

  inline  ::int64_t getId() const;

  inline  ::int64_t getBuildId() const;

  inline  ::int64_t getLayoutId() const;

  inline  ::int64_t getPaintId() const;

  inline bool hasChildren() const;
  inline  ::capnp::List< ::Node>::Reader getChildren() const;

  inline bool hasLayoutGuide() const;
  inline  ::Guide::Reader getLayoutGuide() const;

  inline bool hasValues() const;
  inline  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Reader getValues() const;

private:
  ::capnp::_::StructReader _reader;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::_::PointerHelpers;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::List;
  friend class ::capnp::MessageBuilder;
  friend class ::capnp::Orphanage;
};

class Node::Builder {
public:
  typedef Node Builds;

  Builder() = delete;  // Deleted to discourage incorrect usage.
                       // You can explicitly initialize to nullptr instead.
  inline Builder(decltype(nullptr)) {}
  inline explicit Builder(::capnp::_::StructBuilder base): _builder(base) {}
  inline operator Reader() const { return Reader(_builder.asReader()); }
  inline Reader asReader() const { return *this; }

  inline ::capnp::MessageSize totalSize() const { return asReader().totalSize(); }
#if !CAPNP_LITE
  inline ::kj::StringTree toString() const { return asReader().toString(); }
#endif  // !CAPNP_LITE

  inline  ::int64_t getId();
  inline void setId( ::int64_t value);

  inline  ::int64_t getBuildId();
  inline void setBuildId( ::int64_t value);

  inline  ::int64_t getLayoutId();
  inline void setLayoutId( ::int64_t value);

  inline  ::int64_t getPaintId();
  inline void setPaintId( ::int64_t value);

  inline bool hasChildren();
  inline  ::capnp::List< ::Node>::Builder getChildren();
  inline void setChildren( ::capnp::List< ::Node>::Reader value);
  inline  ::capnp::List< ::Node>::Builder initChildren(unsigned int size);
  inline void adoptChildren(::capnp::Orphan< ::capnp::List< ::Node>>&& value);
  inline ::capnp::Orphan< ::capnp::List< ::Node>> disownChildren();

  inline bool hasLayoutGuide();
  inline  ::Guide::Builder getLayoutGuide();
  inline void setLayoutGuide( ::Guide::Reader value);
  inline  ::Guide::Builder initLayoutGuide();
  inline void adoptLayoutGuide(::capnp::Orphan< ::Guide>&& value);
  inline ::capnp::Orphan< ::Guide> disownLayoutGuide();

  inline bool hasValues();
  inline  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Builder getValues();
  inline void setValues( ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Reader value);
  inline  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Builder initValues();
  inline void adoptValues(::capnp::Orphan< ::Map< ::capnp::Text,  ::capnp::AnyPointer>>&& value);
  inline ::capnp::Orphan< ::Map< ::capnp::Text,  ::capnp::AnyPointer>> disownValues();

private:
  ::capnp::_::StructBuilder _builder;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
  friend class ::capnp::Orphanage;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::_::PointerHelpers;
};

#if !CAPNP_LITE
class Node::Pipeline {
public:
  typedef Node Pipelines;

  inline Pipeline(decltype(nullptr)): _typeless(nullptr) {}
  inline explicit Pipeline(::capnp::AnyPointer::Pipeline&& typeless)
      : _typeless(kj::mv(typeless)) {}

  inline  ::Guide::Pipeline getLayoutGuide();
  inline  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Pipeline getValues();
private:
  ::capnp::AnyPointer::Pipeline _typeless;
  friend class ::capnp::PipelineHook;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
};
#endif  // !CAPNP_LITE

class Root::Reader {
public:
  typedef Root Reads;

  Reader() = default;
  inline explicit Reader(::capnp::_::StructReader base): _reader(base) {}

  inline ::capnp::MessageSize totalSize() const {
    return _reader.totalSize().asPublic();
  }

#if !CAPNP_LITE
  inline ::kj::StringTree toString() const {
    return ::capnp::_::structString(_reader, *_capnpPrivate::brand);
  }
#endif  // !CAPNP_LITE

  inline bool hasNode() const;
  inline  ::Node::Reader getNode() const;

private:
  ::capnp::_::StructReader _reader;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::_::PointerHelpers;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::List;
  friend class ::capnp::MessageBuilder;
  friend class ::capnp::Orphanage;
};

class Root::Builder {
public:
  typedef Root Builds;

  Builder() = delete;  // Deleted to discourage incorrect usage.
                       // You can explicitly initialize to nullptr instead.
  inline Builder(decltype(nullptr)) {}
  inline explicit Builder(::capnp::_::StructBuilder base): _builder(base) {}
  inline operator Reader() const { return Reader(_builder.asReader()); }
  inline Reader asReader() const { return *this; }

  inline ::capnp::MessageSize totalSize() const { return asReader().totalSize(); }
#if !CAPNP_LITE
  inline ::kj::StringTree toString() const { return asReader().toString(); }
#endif  // !CAPNP_LITE

  inline bool hasNode();
  inline  ::Node::Builder getNode();
  inline void setNode( ::Node::Reader value);
  inline  ::Node::Builder initNode();
  inline void adoptNode(::capnp::Orphan< ::Node>&& value);
  inline ::capnp::Orphan< ::Node> disownNode();

private:
  ::capnp::_::StructBuilder _builder;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
  friend class ::capnp::Orphanage;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::_::PointerHelpers;
};

#if !CAPNP_LITE
class Root::Pipeline {
public:
  typedef Root Pipelines;

  inline Pipeline(decltype(nullptr)): _typeless(nullptr) {}
  inline explicit Pipeline(::capnp::AnyPointer::Pipeline&& typeless)
      : _typeless(kj::mv(typeless)) {}

  inline  ::Node::Pipeline getNode();
private:
  ::capnp::AnyPointer::Pipeline _typeless;
  friend class ::capnp::PipelineHook;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
};
#endif  // !CAPNP_LITE

template <typename Key, typename Value>
class Map<Key, Value>::Reader {
public:
  typedef Map Reads;

  Reader() = default;
  inline explicit Reader(::capnp::_::StructReader base): _reader(base) {}

  inline ::capnp::MessageSize totalSize() const {
    return _reader.totalSize().asPublic();
  }

#if !CAPNP_LITE
  inline ::kj::StringTree toString() const {
    return ::capnp::_::structString(_reader, *_capnpPrivate::brand);
  }
#endif  // !CAPNP_LITE

  inline bool hasEntries() const;
  inline typename  ::capnp::List<typename  ::Map<Key, Value>::Entry>::Reader getEntries() const;

private:
  ::capnp::_::StructReader _reader;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::_::PointerHelpers;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::List;
  friend class ::capnp::MessageBuilder;
  friend class ::capnp::Orphanage;
};

template <typename Key, typename Value>
class Map<Key, Value>::Builder {
public:
  typedef Map Builds;

  Builder() = delete;  // Deleted to discourage incorrect usage.
                       // You can explicitly initialize to nullptr instead.
  inline Builder(decltype(nullptr)) {}
  inline explicit Builder(::capnp::_::StructBuilder base): _builder(base) {}
  inline operator Reader() const { return Reader(_builder.asReader()); }
  inline Reader asReader() const { return *this; }

  inline ::capnp::MessageSize totalSize() const { return asReader().totalSize(); }
#if !CAPNP_LITE
  inline ::kj::StringTree toString() const { return asReader().toString(); }
#endif  // !CAPNP_LITE

  inline bool hasEntries();
  inline typename  ::capnp::List<typename  ::Map<Key, Value>::Entry>::Builder getEntries();
  inline void setEntries(typename  ::capnp::List<typename  ::Map<Key, Value>::Entry>::Reader value);
  inline typename  ::capnp::List<typename  ::Map<Key, Value>::Entry>::Builder initEntries(unsigned int size);
  inline void adoptEntries(::capnp::Orphan< ::capnp::List<typename  ::Map<Key, Value>::Entry>>&& value);
  inline ::capnp::Orphan< ::capnp::List<typename  ::Map<Key, Value>::Entry>> disownEntries();

private:
  ::capnp::_::StructBuilder _builder;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
  friend class ::capnp::Orphanage;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::_::PointerHelpers;
};

#if !CAPNP_LITE
template <typename Key, typename Value>
class Map<Key, Value>::Pipeline {
public:
  typedef Map Pipelines;

  inline Pipeline(decltype(nullptr)): _typeless(nullptr) {}
  inline explicit Pipeline(::capnp::AnyPointer::Pipeline&& typeless)
      : _typeless(kj::mv(typeless)) {}

private:
  ::capnp::AnyPointer::Pipeline _typeless;
  friend class ::capnp::PipelineHook;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
};
#endif  // !CAPNP_LITE

template <typename Key, typename Value>
class Map<Key, Value>::Entry::Reader {
public:
  typedef Entry Reads;

  Reader() = default;
  inline explicit Reader(::capnp::_::StructReader base): _reader(base) {}

  inline ::capnp::MessageSize totalSize() const {
    return _reader.totalSize().asPublic();
  }

#if !CAPNP_LITE
  inline ::kj::StringTree toString() const {
    return ::capnp::_::structString(_reader, *_capnpPrivate::brand);
  }
#endif  // !CAPNP_LITE

  inline bool hasKey() const;
  inline  ::capnp::ReaderFor<Key> getKey() const;

  inline bool hasValue() const;
  inline  ::capnp::ReaderFor<Value> getValue() const;

private:
  ::capnp::_::StructReader _reader;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::_::PointerHelpers;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::List;
  friend class ::capnp::MessageBuilder;
  friend class ::capnp::Orphanage;
};

template <typename Key, typename Value>
class Map<Key, Value>::Entry::Builder {
public:
  typedef Entry Builds;

  Builder() = delete;  // Deleted to discourage incorrect usage.
                       // You can explicitly initialize to nullptr instead.
  inline Builder(decltype(nullptr)) {}
  inline explicit Builder(::capnp::_::StructBuilder base): _builder(base) {}
  inline operator Reader() const { return Reader(_builder.asReader()); }
  inline Reader asReader() const { return *this; }

  inline ::capnp::MessageSize totalSize() const { return asReader().totalSize(); }
#if !CAPNP_LITE
  inline ::kj::StringTree toString() const { return asReader().toString(); }
#endif  // !CAPNP_LITE

  inline bool hasKey();
  inline  ::capnp::BuilderFor<Key> getKey();
  inline void setKey( ::capnp::ReaderFor<Key> value);
  inline  ::capnp::BuilderFor<Key> initKey();
  inline  ::capnp::BuilderFor<Key> initKey(unsigned int size);
  inline void adoptKey(::capnp::Orphan<Key>&& value);
  inline ::capnp::Orphan<Key> disownKey();

  inline bool hasValue();
  inline  ::capnp::BuilderFor<Value> getValue();
  inline void setValue( ::capnp::ReaderFor<Value> value);
  inline  ::capnp::BuilderFor<Value> initValue();
  inline  ::capnp::BuilderFor<Value> initValue(unsigned int size);
  inline void adoptValue(::capnp::Orphan<Value>&& value);
  inline ::capnp::Orphan<Value> disownValue();

private:
  ::capnp::_::StructBuilder _builder;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
  friend class ::capnp::Orphanage;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::_::PointerHelpers;
};

#if !CAPNP_LITE
template <typename Key, typename Value>
class Map<Key, Value>::Entry::Pipeline {
public:
  typedef Entry Pipelines;

  inline Pipeline(decltype(nullptr)): _typeless(nullptr) {}
  inline explicit Pipeline(::capnp::AnyPointer::Pipeline&& typeless)
      : _typeless(kj::mv(typeless)) {}

  inline  ::capnp::PipelineFor<Key> getKey();
  inline  ::capnp::PipelineFor<Value> getValue();
private:
  ::capnp::AnyPointer::Pipeline _typeless;
  friend class ::capnp::PipelineHook;
  template <typename, ::capnp::Kind>
  friend struct ::capnp::ToDynamic_;
};
#endif  // !CAPNP_LITE

// =======================================================================================

inline  ::int64_t Node::Reader::getId() const {
  return _reader.getDataField< ::int64_t>(
      0 * ::capnp::ELEMENTS);
}

inline  ::int64_t Node::Builder::getId() {
  return _builder.getDataField< ::int64_t>(
      0 * ::capnp::ELEMENTS);
}
inline void Node::Builder::setId( ::int64_t value) {
  _builder.setDataField< ::int64_t>(
      0 * ::capnp::ELEMENTS, value);
}

inline  ::int64_t Node::Reader::getBuildId() const {
  return _reader.getDataField< ::int64_t>(
      1 * ::capnp::ELEMENTS);
}

inline  ::int64_t Node::Builder::getBuildId() {
  return _builder.getDataField< ::int64_t>(
      1 * ::capnp::ELEMENTS);
}
inline void Node::Builder::setBuildId( ::int64_t value) {
  _builder.setDataField< ::int64_t>(
      1 * ::capnp::ELEMENTS, value);
}

inline  ::int64_t Node::Reader::getLayoutId() const {
  return _reader.getDataField< ::int64_t>(
      2 * ::capnp::ELEMENTS);
}

inline  ::int64_t Node::Builder::getLayoutId() {
  return _builder.getDataField< ::int64_t>(
      2 * ::capnp::ELEMENTS);
}
inline void Node::Builder::setLayoutId( ::int64_t value) {
  _builder.setDataField< ::int64_t>(
      2 * ::capnp::ELEMENTS, value);
}

inline  ::int64_t Node::Reader::getPaintId() const {
  return _reader.getDataField< ::int64_t>(
      3 * ::capnp::ELEMENTS);
}

inline  ::int64_t Node::Builder::getPaintId() {
  return _builder.getDataField< ::int64_t>(
      3 * ::capnp::ELEMENTS);
}
inline void Node::Builder::setPaintId( ::int64_t value) {
  _builder.setDataField< ::int64_t>(
      3 * ::capnp::ELEMENTS, value);
}

inline bool Node::Reader::hasChildren() const {
  return !_reader.getPointerField(0 * ::capnp::POINTERS).isNull();
}
inline bool Node::Builder::hasChildren() {
  return !_builder.getPointerField(0 * ::capnp::POINTERS).isNull();
}
inline  ::capnp::List< ::Node>::Reader Node::Reader::getChildren() const {
  return ::capnp::_::PointerHelpers< ::capnp::List< ::Node>>::get(
      _reader.getPointerField(0 * ::capnp::POINTERS));
}
inline  ::capnp::List< ::Node>::Builder Node::Builder::getChildren() {
  return ::capnp::_::PointerHelpers< ::capnp::List< ::Node>>::get(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}
inline void Node::Builder::setChildren( ::capnp::List< ::Node>::Reader value) {
  ::capnp::_::PointerHelpers< ::capnp::List< ::Node>>::set(
      _builder.getPointerField(0 * ::capnp::POINTERS), value);
}
inline  ::capnp::List< ::Node>::Builder Node::Builder::initChildren(unsigned int size) {
  return ::capnp::_::PointerHelpers< ::capnp::List< ::Node>>::init(
      _builder.getPointerField(0 * ::capnp::POINTERS), size);
}
inline void Node::Builder::adoptChildren(
    ::capnp::Orphan< ::capnp::List< ::Node>>&& value) {
  ::capnp::_::PointerHelpers< ::capnp::List< ::Node>>::adopt(
      _builder.getPointerField(0 * ::capnp::POINTERS), kj::mv(value));
}
inline ::capnp::Orphan< ::capnp::List< ::Node>> Node::Builder::disownChildren() {
  return ::capnp::_::PointerHelpers< ::capnp::List< ::Node>>::disown(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}

inline bool Node::Reader::hasLayoutGuide() const {
  return !_reader.getPointerField(1 * ::capnp::POINTERS).isNull();
}
inline bool Node::Builder::hasLayoutGuide() {
  return !_builder.getPointerField(1 * ::capnp::POINTERS).isNull();
}
inline  ::Guide::Reader Node::Reader::getLayoutGuide() const {
  return ::capnp::_::PointerHelpers< ::Guide>::get(
      _reader.getPointerField(1 * ::capnp::POINTERS));
}
inline  ::Guide::Builder Node::Builder::getLayoutGuide() {
  return ::capnp::_::PointerHelpers< ::Guide>::get(
      _builder.getPointerField(1 * ::capnp::POINTERS));
}
#if !CAPNP_LITE
inline  ::Guide::Pipeline Node::Pipeline::getLayoutGuide() {
  return  ::Guide::Pipeline(_typeless.getPointerField(1));
}
#endif  // !CAPNP_LITE
inline void Node::Builder::setLayoutGuide( ::Guide::Reader value) {
  ::capnp::_::PointerHelpers< ::Guide>::set(
      _builder.getPointerField(1 * ::capnp::POINTERS), value);
}
inline  ::Guide::Builder Node::Builder::initLayoutGuide() {
  return ::capnp::_::PointerHelpers< ::Guide>::init(
      _builder.getPointerField(1 * ::capnp::POINTERS));
}
inline void Node::Builder::adoptLayoutGuide(
    ::capnp::Orphan< ::Guide>&& value) {
  ::capnp::_::PointerHelpers< ::Guide>::adopt(
      _builder.getPointerField(1 * ::capnp::POINTERS), kj::mv(value));
}
inline ::capnp::Orphan< ::Guide> Node::Builder::disownLayoutGuide() {
  return ::capnp::_::PointerHelpers< ::Guide>::disown(
      _builder.getPointerField(1 * ::capnp::POINTERS));
}

inline bool Node::Reader::hasValues() const {
  return !_reader.getPointerField(2 * ::capnp::POINTERS).isNull();
}
inline bool Node::Builder::hasValues() {
  return !_builder.getPointerField(2 * ::capnp::POINTERS).isNull();
}
inline  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Reader Node::Reader::getValues() const {
  return ::capnp::_::PointerHelpers< ::Map< ::capnp::Text,  ::capnp::AnyPointer>>::get(
      _reader.getPointerField(2 * ::capnp::POINTERS));
}
inline  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Builder Node::Builder::getValues() {
  return ::capnp::_::PointerHelpers< ::Map< ::capnp::Text,  ::capnp::AnyPointer>>::get(
      _builder.getPointerField(2 * ::capnp::POINTERS));
}
#if !CAPNP_LITE
inline  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Pipeline Node::Pipeline::getValues() {
  return  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Pipeline(_typeless.getPointerField(2));
}
#endif  // !CAPNP_LITE
inline void Node::Builder::setValues( ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Reader value) {
  ::capnp::_::PointerHelpers< ::Map< ::capnp::Text,  ::capnp::AnyPointer>>::set(
      _builder.getPointerField(2 * ::capnp::POINTERS), value);
}
inline  ::Map< ::capnp::Text,  ::capnp::AnyPointer>::Builder Node::Builder::initValues() {
  return ::capnp::_::PointerHelpers< ::Map< ::capnp::Text,  ::capnp::AnyPointer>>::init(
      _builder.getPointerField(2 * ::capnp::POINTERS));
}
inline void Node::Builder::adoptValues(
    ::capnp::Orphan< ::Map< ::capnp::Text,  ::capnp::AnyPointer>>&& value) {
  ::capnp::_::PointerHelpers< ::Map< ::capnp::Text,  ::capnp::AnyPointer>>::adopt(
      _builder.getPointerField(2 * ::capnp::POINTERS), kj::mv(value));
}
inline ::capnp::Orphan< ::Map< ::capnp::Text,  ::capnp::AnyPointer>> Node::Builder::disownValues() {
  return ::capnp::_::PointerHelpers< ::Map< ::capnp::Text,  ::capnp::AnyPointer>>::disown(
      _builder.getPointerField(2 * ::capnp::POINTERS));
}

inline bool Root::Reader::hasNode() const {
  return !_reader.getPointerField(0 * ::capnp::POINTERS).isNull();
}
inline bool Root::Builder::hasNode() {
  return !_builder.getPointerField(0 * ::capnp::POINTERS).isNull();
}
inline  ::Node::Reader Root::Reader::getNode() const {
  return ::capnp::_::PointerHelpers< ::Node>::get(
      _reader.getPointerField(0 * ::capnp::POINTERS));
}
inline  ::Node::Builder Root::Builder::getNode() {
  return ::capnp::_::PointerHelpers< ::Node>::get(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}
#if !CAPNP_LITE
inline  ::Node::Pipeline Root::Pipeline::getNode() {
  return  ::Node::Pipeline(_typeless.getPointerField(0));
}
#endif  // !CAPNP_LITE
inline void Root::Builder::setNode( ::Node::Reader value) {
  ::capnp::_::PointerHelpers< ::Node>::set(
      _builder.getPointerField(0 * ::capnp::POINTERS), value);
}
inline  ::Node::Builder Root::Builder::initNode() {
  return ::capnp::_::PointerHelpers< ::Node>::init(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}
inline void Root::Builder::adoptNode(
    ::capnp::Orphan< ::Node>&& value) {
  ::capnp::_::PointerHelpers< ::Node>::adopt(
      _builder.getPointerField(0 * ::capnp::POINTERS), kj::mv(value));
}
inline ::capnp::Orphan< ::Node> Root::Builder::disownNode() {
  return ::capnp::_::PointerHelpers< ::Node>::disown(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}

template <typename Key, typename Value>
inline bool Map<Key, Value>::Reader::hasEntries() const {
  return !_reader.getPointerField(0 * ::capnp::POINTERS).isNull();
}
template <typename Key, typename Value>
inline bool Map<Key, Value>::Builder::hasEntries() {
  return !_builder.getPointerField(0 * ::capnp::POINTERS).isNull();
}
template <typename Key, typename Value>
inline typename  ::capnp::List<typename  ::Map<Key, Value>::Entry>::Reader Map<Key, Value>::Reader::getEntries() const {
  return ::capnp::_::PointerHelpers< ::capnp::List<typename  ::Map<Key, Value>::Entry>>::get(
      _reader.getPointerField(0 * ::capnp::POINTERS));
}
template <typename Key, typename Value>
inline typename  ::capnp::List<typename  ::Map<Key, Value>::Entry>::Builder Map<Key, Value>::Builder::getEntries() {
  return ::capnp::_::PointerHelpers< ::capnp::List<typename  ::Map<Key, Value>::Entry>>::get(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}
template <typename Key, typename Value>
inline void Map<Key, Value>::Builder::setEntries(typename  ::capnp::List<typename  ::Map<Key, Value>::Entry>::Reader value) {
  ::capnp::_::PointerHelpers< ::capnp::List<typename  ::Map<Key, Value>::Entry>>::set(
      _builder.getPointerField(0 * ::capnp::POINTERS), value);
}
template <typename Key, typename Value>
inline typename  ::capnp::List<typename  ::Map<Key, Value>::Entry>::Builder Map<Key, Value>::Builder::initEntries(unsigned int size) {
  return ::capnp::_::PointerHelpers< ::capnp::List<typename  ::Map<Key, Value>::Entry>>::init(
      _builder.getPointerField(0 * ::capnp::POINTERS), size);
}
template <typename Key, typename Value>
inline void Map<Key, Value>::Builder::adoptEntries(
    ::capnp::Orphan< ::capnp::List<typename  ::Map<Key, Value>::Entry>>&& value) {
  ::capnp::_::PointerHelpers< ::capnp::List<typename  ::Map<Key, Value>::Entry>>::adopt(
      _builder.getPointerField(0 * ::capnp::POINTERS), kj::mv(value));
}
template <typename Key, typename Value>
inline ::capnp::Orphan< ::capnp::List<typename  ::Map<Key, Value>::Entry>> Map<Key, Value>::Builder::disownEntries() {
  return ::capnp::_::PointerHelpers< ::capnp::List<typename  ::Map<Key, Value>::Entry>>::disown(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}

template <typename Key, typename Value>
inline bool Map<Key, Value>::Entry::Reader::hasKey() const {
  return !_reader.getPointerField(0 * ::capnp::POINTERS).isNull();
}
template <typename Key, typename Value>
inline bool Map<Key, Value>::Entry::Builder::hasKey() {
  return !_builder.getPointerField(0 * ::capnp::POINTERS).isNull();
}
template <typename Key, typename Value>
inline  ::capnp::ReaderFor<Key> Map<Key, Value>::Entry::Reader::getKey() const {
  return ::capnp::_::PointerHelpers<Key>::get(
      _reader.getPointerField(0 * ::capnp::POINTERS));
}
template <typename Key, typename Value>
inline  ::capnp::BuilderFor<Key> Map<Key, Value>::Entry::Builder::getKey() {
  return ::capnp::_::PointerHelpers<Key>::get(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}
#if !CAPNP_LITE
template <typename Key, typename Value>
inline  ::capnp::PipelineFor<Key> Map<Key, Value>::Entry::Pipeline::getKey() {
  return  ::capnp::PipelineFor<Key>(_typeless.getPointerField(0));
}
#endif  // !CAPNP_LITE
template <typename Key, typename Value>
inline void Map<Key, Value>::Entry::Builder::setKey( ::capnp::ReaderFor<Key> value) {
  ::capnp::_::PointerHelpers<Key>::set(
      _builder.getPointerField(0 * ::capnp::POINTERS), value);
}
template <typename Key, typename Value>
inline  ::capnp::BuilderFor<Key> Map<Key, Value>::Entry::Builder::initKey() {
  return ::capnp::_::PointerHelpers<Key>::init(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}
template <typename Key, typename Value>
inline  ::capnp::BuilderFor<Key> Map<Key, Value>::Entry::Builder::initKey(unsigned int size) {
  return ::capnp::_::PointerHelpers<Key>::init(
      _builder.getPointerField(0 * ::capnp::POINTERS), size);
}
template <typename Key, typename Value>
inline void Map<Key, Value>::Entry::Builder::adoptKey(
    ::capnp::Orphan<Key>&& value) {
  ::capnp::_::PointerHelpers<Key>::adopt(
      _builder.getPointerField(0 * ::capnp::POINTERS), kj::mv(value));
}
template <typename Key, typename Value>
inline ::capnp::Orphan<Key> Map<Key, Value>::Entry::Builder::disownKey() {
  return ::capnp::_::PointerHelpers<Key>::disown(
      _builder.getPointerField(0 * ::capnp::POINTERS));
}

template <typename Key, typename Value>
inline bool Map<Key, Value>::Entry::Reader::hasValue() const {
  return !_reader.getPointerField(1 * ::capnp::POINTERS).isNull();
}
template <typename Key, typename Value>
inline bool Map<Key, Value>::Entry::Builder::hasValue() {
  return !_builder.getPointerField(1 * ::capnp::POINTERS).isNull();
}
template <typename Key, typename Value>
inline  ::capnp::ReaderFor<Value> Map<Key, Value>::Entry::Reader::getValue() const {
  return ::capnp::_::PointerHelpers<Value>::get(
      _reader.getPointerField(1 * ::capnp::POINTERS));
}
template <typename Key, typename Value>
inline  ::capnp::BuilderFor<Value> Map<Key, Value>::Entry::Builder::getValue() {
  return ::capnp::_::PointerHelpers<Value>::get(
      _builder.getPointerField(1 * ::capnp::POINTERS));
}
#if !CAPNP_LITE
template <typename Key, typename Value>
inline  ::capnp::PipelineFor<Value> Map<Key, Value>::Entry::Pipeline::getValue() {
  return  ::capnp::PipelineFor<Value>(_typeless.getPointerField(1));
}
#endif  // !CAPNP_LITE
template <typename Key, typename Value>
inline void Map<Key, Value>::Entry::Builder::setValue( ::capnp::ReaderFor<Value> value) {
  ::capnp::_::PointerHelpers<Value>::set(
      _builder.getPointerField(1 * ::capnp::POINTERS), value);
}
template <typename Key, typename Value>
inline  ::capnp::BuilderFor<Value> Map<Key, Value>::Entry::Builder::initValue() {
  return ::capnp::_::PointerHelpers<Value>::init(
      _builder.getPointerField(1 * ::capnp::POINTERS));
}
template <typename Key, typename Value>
inline  ::capnp::BuilderFor<Value> Map<Key, Value>::Entry::Builder::initValue(unsigned int size) {
  return ::capnp::_::PointerHelpers<Value>::init(
      _builder.getPointerField(1 * ::capnp::POINTERS), size);
}
template <typename Key, typename Value>
inline void Map<Key, Value>::Entry::Builder::adoptValue(
    ::capnp::Orphan<Value>&& value) {
  ::capnp::_::PointerHelpers<Value>::adopt(
      _builder.getPointerField(1 * ::capnp::POINTERS), kj::mv(value));
}
template <typename Key, typename Value>
inline ::capnp::Orphan<Value> Map<Key, Value>::Entry::Builder::disownValue() {
  return ::capnp::_::PointerHelpers<Value>::disown(
      _builder.getPointerField(1 * ::capnp::POINTERS));
}

// Map<Key, Value>::Entry
#ifndef _MSC_VER
template <typename Key, typename Value>
constexpr uint16_t Map<Key, Value>::Entry::_capnpPrivate::dataWordSize;
template <typename Key, typename Value>
constexpr uint16_t Map<Key, Value>::Entry::_capnpPrivate::pointerCount;
#endif
#if !CAPNP_LITE
template <typename Key, typename Value>
constexpr ::capnp::Kind Map<Key, Value>::Entry::_capnpPrivate::kind;
template <typename Key, typename Value>
constexpr ::capnp::_::RawSchema const* Map<Key, Value>::Entry::_capnpPrivate::schema;
template <typename Key, typename Value>
constexpr ::capnp::_::RawBrandedSchema const* Map<Key, Value>::Entry::_capnpPrivate::brand;
template <typename Key, typename Value>
const ::capnp::_::RawBrandedSchema::Scope Map<Key, Value>::Entry::_capnpPrivate::brandScopes[] = {
  { 0xbd0993cece143673, brandBindings + 0, 2, false},
};
template <typename Key, typename Value>
const ::capnp::_::RawBrandedSchema::Binding Map<Key, Value>::Entry::_capnpPrivate::brandBindings[] = {
  ::capnp::_::brandBindingFor<Key>(),
  ::capnp::_::brandBindingFor<Value>(),
};
template <typename Key, typename Value>
const ::capnp::_::RawBrandedSchema Map<Key, Value>::Entry::_capnpPrivate::specificBrand = {
  &::capnp::schemas::s_91412d97d3671d43, brandScopes, nullptr,
  sizeof(brandScopes) / sizeof(brandScopes[0]), 0, nullptr
};
#endif  // !CAPNP_LITE

// Map<Key, Value>
#ifndef _MSC_VER
template <typename Key, typename Value>
constexpr uint16_t Map<Key, Value>::_capnpPrivate::dataWordSize;
template <typename Key, typename Value>
constexpr uint16_t Map<Key, Value>::_capnpPrivate::pointerCount;
#endif
#if !CAPNP_LITE
template <typename Key, typename Value>
constexpr ::capnp::Kind Map<Key, Value>::_capnpPrivate::kind;
template <typename Key, typename Value>
constexpr ::capnp::_::RawSchema const* Map<Key, Value>::_capnpPrivate::schema;
template <typename Key, typename Value>
constexpr ::capnp::_::RawBrandedSchema const* Map<Key, Value>::_capnpPrivate::brand;
template <typename Key, typename Value>
const ::capnp::_::RawBrandedSchema::Scope Map<Key, Value>::_capnpPrivate::brandScopes[] = {
  { 0xbd0993cece143673, brandBindings + 0, 2, false},
};
template <typename Key, typename Value>
const ::capnp::_::RawBrandedSchema::Binding Map<Key, Value>::_capnpPrivate::brandBindings[] = {
  ::capnp::_::brandBindingFor<Key>(),
  ::capnp::_::brandBindingFor<Value>(),
};
template <typename Key, typename Value>
const ::capnp::_::RawBrandedSchema::Dependency Map<Key, Value>::_capnpPrivate::brandDependencies[] = {
  { 16777216,  ::Map<Key, Value>::Entry::_capnpPrivate::brand },
};
template <typename Key, typename Value>
const ::capnp::_::RawBrandedSchema Map<Key, Value>::_capnpPrivate::specificBrand = {
  &::capnp::schemas::s_bd0993cece143673, brandScopes, brandDependencies,
  sizeof(brandScopes) / sizeof(brandScopes[0]), sizeof(brandDependencies) / sizeof(brandDependencies[0]), nullptr
};
#endif  // !CAPNP_LITE


#endif  // CAPNP_INCLUDED_9abec47167560884_
