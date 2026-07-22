interface ProductsInterface {
    void getProductVariants();
}

class AlgoliaProducts implements ProductsInterface {
    @Override
    public void getProductVariants() {
        System.out.println("Algolia StoreProducts");
    }
}

class AtlasProducts {
    public void getAtlasProductVariants() {
        System.out.println("Atlas StoreProducts");
    }
}

class AtlasProductsAdapter implements ProductsInterface {
    private final AtlasProducts atlasProducts;

    public AtlasProductsAdapter(AtlasProducts atlasProducts) {
        this.atlasProducts = atlasProducts;
    }

    @Override
    public void getProductVariants() {
        atlasProducts.getAtlasProductVariants();
    }
}

public class AdapterPattern {
    public static void main(String[] args) {
        ProductsInterface algolia = new AlgoliaProducts();
        algolia.getProductVariants();

        AtlasProducts atlas = new AtlasProducts();
        ProductsInterface atlasAdapter = new AtlasProductsAdapter(atlas);
        atlasAdapter.getProductVariants();
    }
}
